### Simple Convert Response to CSV

Convert response repository to csv and set header text/csv in handler layer

#### service.go
```go
func (college *service) GetUserCSV() ([]byte, error) {
  data, err := college.Get()
  if err != nil {
    return nil, err
}

csvbuffer:= bytes.Buffer{} //data is stored in the buffer
writer := csv.NewWriter(&csvbuffer)

header := []string{"Nim", "Name", "Campus"} //first row for the header
  if err := writer.Write(header); err != nil {
    return nil, err
}

for _, v := range data {
  var datacsv []string
  datacsv = append(datacsv, v.Nim, v.Name, v.Campus)
  if err := writer.Write(datacsv); err != nil {
    return nil, err
  }
}

writer.Flush() //Ensure all data has been written to the buffer

if err := writer.Error(); err != nil {
  return nil, err
}

return csvbuffer.Bytes(), nil
}
```

handler.go
```go
func (college *handler) CsvUser(f *fiber.Ctx)error {
  result,err := college.service.GetUserCSV()
  if err != nil {
    return f.Status(400).JSON(fiber.Map{
      "message": err.Error(),
    })
  }

  f.Set(fiber.HeaderContentType, "text/csv")
  f.Set(fiber.HeaderContentDisposition, "attachment; filename=datauser.csv")

  return f.Send(result)
}
```
