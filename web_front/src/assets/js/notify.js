export function notifyApiResult(instance, returnCode, msg) {
  if (returnCode !== 0) {
    instance.$message({
      title: 'Result',
      type: 'error',
      message: msg
    })
  } else {
    instance.$message({
      title: 'Result',
      type: 'success',
      message: '调用成功'
    })
  }
}
