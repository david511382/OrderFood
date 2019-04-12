var submitButton = $( "#SubmitButton" );
submitButton.click(send);

$.ajax({
    type:"GET",
    url: "/api/order/all"
}).done(showTotalOrders);

function send(event){
    event.preventDefault();    
    
    var msg = $("#shopcart").text();
    var data ={orders: msg};
    $.ajax({
        type:"PUT",
        url: "/api/order",  
        data:data
    }).done(showTotalOrders);

    alert("送出\n"+ msg);
}

function showTotalOrders(data){
    $("#result").html(data);
}

