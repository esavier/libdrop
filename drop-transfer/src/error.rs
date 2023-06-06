use std::io::Error as IoError;

use tokio_tungstenite::tungstenite;

#[derive(thiserror::Error, Debug)]
pub enum Error {
    #[error("Operation was canceled")]
    Canceled,
    #[error("Invalid path: {0}")]
    BadPath(String),
    #[error("Could not open file")]
    BadFile,
    #[error("Service failed to stop")]
    ServiceStop,
    #[error("Transfer not found")]
    BadTransfer,
    #[error("Invalid transfer state: {0}")]
    BadTransferState(String),
    #[error("Invalid file ID")]
    BadFileId,
    #[error("File size has changed")]
    MismatchedSize,
    #[error("Unexpected data")]
    UnexpectedData,
    #[error("IO error {0}")]
    Io(#[from] IoError),
    #[error("Got directory when expecting a file")]
    DirectoryNotExpected,
    #[error("Transfer limits exceeded")]
    TransferLimitsExceeded,
    #[error("Invalid argument")]
    InvalidArgument,
    #[error("Transfer timeout")]
    TransferTimeout,
    #[error("Server connection failure: {0}")]
    WsServer(#[from] warp::Error),
    #[error("Client connection failure: {0}")]
    WsClient(#[from] tungstenite::Error),
    #[error("Address already in use")]
    AddrInUse,
    #[error("File modified")]
    FileModified,
    #[error("Filename is too long")]
    FilenameTooLong,
    #[error("Failed to authenticate to the peer")]
    AuthenticationFailed,
    #[error("Storage error")]
    StorageError,
}

impl Error {
    pub fn os_err_code(&self) -> Option<i32> {
        match self {
            Error::Io(ioerr) => ioerr.raw_os_error().map(|c| c as _),
            Error::WsServer(_) => {
                // TODO(msz): Theoretically it should be possible to extract OS error from WS
                // server but warp does not make it easy. Maybe one
                // day we will rewrite warp server with raw tungstenite
                None
            }
            Error::WsClient(terr) => {
                if let tungstenite::Error::Io(ioerr) = terr {
                    ioerr.raw_os_error().map(|c| c as _)
                } else {
                    None
                }
            }
            _ => None,
        }
    }
}

impl From<&Error> for u32 {
    fn from(err: &Error) -> Self {
        match err {
            Error::Canceled => 1,
            Error::BadPath(_) => 2,
            Error::BadFile => 3,
            Error::ServiceStop => 6,
            Error::BadTransfer => 7,
            Error::BadTransferState(_) => 8,
            Error::BadFileId => 9,
            Error::Io(_) => 15,
            Error::DirectoryNotExpected => 17,
            Error::TransferLimitsExceeded => 20,
            Error::MismatchedSize => 21,
            Error::UnexpectedData => 22,
            Error::InvalidArgument => 23,
            Error::TransferTimeout => 24,
            Error::WsServer(_) => 25,
            Error::WsClient(_) => 26,
            Error::AddrInUse => 27,
            Error::FileModified => 28,
            Error::FilenameTooLong => 29,
            Error::AuthenticationFailed => 30,
            Error::StorageError => 31,
        }
    }
}

pub trait ResultExt {
    fn to_status(&self) -> Result<(), i32>;
}

impl<T> ResultExt for super::Result<T> {
    fn to_status(&self) -> Result<(), i32> {
        match self {
            Ok(_) => Ok(()),
            Err(err) => Err(u32::from(err) as _),
        }
    }
}

impl From<walkdir::Error> for Error {
    fn from(value: walkdir::Error) -> Self {
        value
            .into_io_error()
            .map(Into::into)
            .unwrap_or(Error::BadPath("Filesystem loop detected".into()))
    }
}
