function send(event){
    event.preventDefault();    
    var msg = $("#shopcart").text();
    var data ={orders: msg};
    $.ajax({
        type:"POST",
        url: "/post/order",  
        data:data
    }).done(showTotalOrders);

    alert("送出\n"+ msg);
}

$( "#orderForm" ).submit(send);

function showTotalOrders(data){
    $("#result").html(data);
}

$.ajax({
    type:"POST",
    url: "/get/order"
}).done(showTotalOrders);