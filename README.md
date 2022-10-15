# Overview
PTZ camera emulator controlled over HTTP. It streams MJPEG video over HTTP.
Just for fun and to learn Golang :)

![](/media/readme_img_1.png?raw=true)

# REST API
- Set preset
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"Pan": 200,"Tilt": 200, "Zoom": 10}' \
  http://localhost:8080/setPreset
```


