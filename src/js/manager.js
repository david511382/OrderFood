
ws.onmessage = reciveWsHandler;  

var form = $( "#selectViewForm" );
var viewSelect = $( "#viewSelect" );
form.submit(changeView);
viewSelect.bind("change",changeView);

getUserOrders();

function reciveWsHandler(evt){
    getUserOrders();

    showTotalOrders(evt.data);
}

function getUserOrders(){
    $.ajax({
        type:"POST",
        url: "/get/user/orders",  
    }).done(showUserOrders);
}

function showUserOrders(data){
    var textarea = document.getElementById("userOrders");
    textarea.innerHTML = data;
}

function changeView(event){
    if (event !==undefined)
        event.preventDefault();    
    
    var data = {view:viewSelect.val()};
    $.ajax({
        type:"POST",
        url: "/post/view",  
        data:data
    }).done(function(data){
        alert("修改為"+ data);
    });
}

