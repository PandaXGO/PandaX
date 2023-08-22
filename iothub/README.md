
## 原始数据上报格式
```text
010304026C00883BF0
```
iothub将原始数据序列化成，在规则链中 msg就是以下数据。
```json
{
  "rowdata": "010304026C00883BF0"
}
```
需要在规则链中通过函数键 进行解析脚本
```js
/*直连设备：tempVal是产品物模型中所定义属性的标识符*/
var tempVal = msg.rowdata;
/*物模型温度标识符*/
msg.temperature = (parseInt('0x'+tempVal.substr(10, 4))*0.1).toFixed(2);
/*物模型湿度标识符*/
msg.humidity = (parseInt('0x'+tempVal.substr(6, 4))*0.1).toFixed(2);
return {msg: msg, metadata: metadata, msgType: msgType};
```
## 属性上报格式
```json
{
   "attribute1": "value1",
   "attribute2": 0
}
```

## 遥测上报格式

```json
{
  "ts": 1689837909000,
  "values": {
    "telemetry1": "value1",
    "telemetry2": 0
  }
}
```
如果边缘无法获取时间
```json
{
    "telemetry1": "value1",
    "telemetry2": 0
}
```

## 网关子设备属性上报格式
devA 为设备ID
```json
{
  "devA": {
    "attribute1": "value1",
    "attribute2": 0
  },
  "devB": {
    "attribute1": "value1",
    "attribute2": 0
  }
}
```

## 网关子设备遥测上报格式
devA 为设备ID
```json
{
  "devA": [
    {
      "ts": 1689837909000,
      "values": {
        "telemetry1": "value1",
        "telemetry2": 0
      }
    }
  ]
}
```

## 网关子设备连接或断开上报格式
```json
{
  "devA": "online",
  "devB": "offline"
}
```
## 命令下发格式
下发的ID和相应的ID相同
```json
{
   "id": "2343",
   "cmd": "restart",
   "content": {
      "firmware_address": "http://xxx.yyy.com",
      "version": "latest",
      "secret": "****",
      "http_method": "GET"
   }
}
```

## 命令响应的格式
success 返回结果必传 content代表输出参数可选
```json
{
  "id": "2343",
  "success": true,
  "content": {
     "aa": "2"
   }
}
```
