var ws = new WebSocket("ws://"+ document.location.host+"/ws");  

//接收到消息时触发  
ws.onmessage = reciveWs;  

function reciveWs(evt){
    showTotalOrders(evt.data);
}