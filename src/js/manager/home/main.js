
ws.onmessage = reciveWsHandler;  

var form = $( "#selectViewForm" );
var viewSelect = $( "#viewSelect" );
form.submit(changeView);
viewSelect.bind("change",changeView);

GetUserOrders(showUserOrders);

function reciveWsHandler(evt){
    GetUserOrders(showUserOrders);

    showTotalOrders(evt.data);
}

function showUserOrders(data){
    var textarea = document.getElementById("userOrders");
    textarea.innerHTML = data;
}

function changeView(event){
    if (event !==undefined)
        event.preventDefault();    
    
    var shopID = viewSelect.val();
    ChangeView(shopID,function(data){
        alert("修改為"+ data);
    });
}

