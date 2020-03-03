## 压缩文件的读写
#### 压缩文件写入步骤
+ 创建并打开文件：os.open(filename)h或os.openfile()得到文件句柄file
+ new一个write := tar.NewWriter(file)
+得到info信息： fileInfo, err := os.Stat(fileName)
+ 或取头部信息：hdr, err := tar.FileInfoHeader(fileInfo, "")
+ 设置 hdr.Size = int64(len(insertByte))
+ 把头部信息写到write： err = write.WriteHeader(hdr)
+ 把要写入的信息写入 write：ret, err := write.Write(insertByte)
+ err = write.Flush()
+ err = write.Close()
#### 压缩文件读出步骤
+ 打开文件：os.open(filename)
+new一个 read := tar.NewReader(file)
+	读取头部信息：hdr, err := read.Next()
+	创建大小合适的空间：var getByte = make([]byte, hdr.Size)

+读出文件数据：	_, err = read.Read(getByte)
