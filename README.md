# Go QR

This is a small webserver that will accept a query parameter `url` on the root route. It uses the `url` to create a QR-Code to display. I'm currently using this to send URLs from my laptop to my phone.

### Running the code

```
go run .
```

### Try some app url schemes

```
localhost:3333?url=calshow://
```
