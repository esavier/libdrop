/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.2
 *
 * Do not make changes to this file unless you know what you are doing--modify
 * the SWIG interface file instead.
 * ----------------------------------------------------------------------------- */

package com.nordsec.norddrop;

public class NordDrop {
  private transient long swigCPtr;
  protected transient boolean swigCMemOwn;

  protected NordDrop(long cPtr, boolean cMemoryOwn) {
    swigCMemOwn = cMemoryOwn;
    swigCPtr = cPtr;
  }

  protected static long getCPtr(NordDrop obj) {
    return (obj == null) ? 0 : obj.swigCPtr;
  }

  @SuppressWarnings("deprecation")
  protected void finalize() {
    delete();
  }

  public synchronized void delete() {
    if (swigCPtr != 0) {
      if (swigCMemOwn) {
        swigCMemOwn = false;
        libnorddropJNI.delete_NordDrop(swigCPtr);
      }
      swigCPtr = 0;
    }
  }

  public NordDrop(INordDropEventCb events, NorddropLogLevel level, INordDropLoggerCb logger, INordDropPubkeyCb pubkeyCb, byte[] privkey, INordDropFdCb fdCb) {
    this(libnorddropJNI.new_NordDrop__SWIG_0(events, level.swigValue(), logger, pubkeyCb, privkey, fdCb), true);
  }

  public NordDrop(INordDropEventCb events, NorddropLogLevel level, INordDropLoggerCb logger, INordDropPubkeyCb pubkeyCb, byte[] privkey) {
    this(libnorddropJNI.new_NordDrop__SWIG_1(events, level.swigValue(), logger, pubkeyCb, privkey), true);
  }

  public NorddropResult start(String listenAddr, String configJson) {
    return NorddropResult.swigToEnum(libnorddropJNI.NordDrop_start(swigCPtr, this, listenAddr, configJson));
  }

  public NorddropResult stop() {
    return NorddropResult.swigToEnum(libnorddropJNI.NordDrop_stop(swigCPtr, this));
  }

  public NorddropResult cancelTransfer(String txid) {
    return NorddropResult.swigToEnum(libnorddropJNI.NordDrop_cancelTransfer(swigCPtr, this, txid));
  }

  public NorddropResult cancelFile(String txid, String fid) {
    return NorddropResult.swigToEnum(libnorddropJNI.NordDrop_cancelFile(swigCPtr, this, txid, fid));
  }

  public NorddropResult rejectFile(String txid, String fid) {
    return NorddropResult.swigToEnum(libnorddropJNI.NordDrop_rejectFile(swigCPtr, this, txid, fid));
  }

  public NorddropResult download(String txid, String fid, String dstPath) {
    return NorddropResult.swigToEnum(libnorddropJNI.NordDrop_download(swigCPtr, this, txid, fid, dstPath));
  }

  public String newTransfer(String peer, String descriptors) {
    return libnorddropJNI.NordDrop_newTransfer(swigCPtr, this, peer, descriptors);
  }

  public NorddropResult purgeTransfers(String txids) {
    return NorddropResult.swigToEnum(libnorddropJNI.NordDrop_purgeTransfers(swigCPtr, this, txids));
  }

  public NorddropResult purgeTransfersUntil(long untilTimestamp) {
    return NorddropResult.swigToEnum(libnorddropJNI.NordDrop_purgeTransfersUntil(swigCPtr, this, untilTimestamp));
  }

  public NorddropResult removeTransferFile(String txid, String fid) {
    return NorddropResult.swigToEnum(libnorddropJNI.NordDrop_removeTransferFile(swigCPtr, this, txid, fid));
  }

  public String getTransfersSince(long sinceTimestamp) {
    return libnorddropJNI.NordDrop_getTransfersSince(swigCPtr, this, sinceTimestamp);
  }

  public static String version() {
    return libnorddropJNI.NordDrop_version();
  }

}
