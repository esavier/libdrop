#[derive(thiserror::Error, Debug)]
pub enum Error {
    #[error("Storage error: {0}")]
    InternalError(String),
    #[error("DB Error: {0}")]
    DBError(#[from] rusqlite::Error),
    #[error("DB Connection Error: {0}")]
    DBConnectionError(#[from] r2d2::Error),
    #[error("URI parsing: {0}")]
    UriParsing(#[from] url::ParseError),
    #[error("Invalid URI scheme: {0}")]
    InvalidUri(String),
}
