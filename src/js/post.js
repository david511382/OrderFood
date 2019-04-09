var submitButton = $( "#SubmitButton" );
submitButton.click(send);

$.ajax({
    type:"POST",
    url: "/get/order"
}).done(showTotalOrders);

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

function showTotalOrders(data){
    $("#result").html(data);
}

