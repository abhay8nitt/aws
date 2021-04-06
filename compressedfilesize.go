
func CompressedSize(filename string) int{
  
  ## Determine uncompressed size of GZIP file
  f, err := os.Open(filename)
  info,_ :=f.Stat()
  compressedSize := info.Size()
  
  r4 := bufio.NewReader(f)
  _size:= int(info.Size())
  x,_:=r4.Discard(_size-4)
  b4,_:=r4.ReadByte()
  b3,_:=r4.ReadByte()
  b2,_:=r4.ReadByte()
  b1,_:=r4.ReadByte()

  uncompressedSize:= (int(b1) << 24) | (int(b2) << 16) + (int(b3) << 8) + int(b4)
  return compressedSize, uncompressedSize
}

func HumanReadable(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "KMGTPE"[exp])
}
