var submitButton = $( "#SubmitButton" );
submitButton.click(send);

$.ajax({
    type:"GET",
    url: "/order/all"
}).done(showTotalOrders);

function send(event){
    event.preventDefault();    
    
    var msg = $("#shopcart").text();
    var data ={orders: msg};
    $.ajax({
        type:"PUT",
        url: "/order",  
        data:data
    }).done(showTotalOrders);

    alert("送出\n"+ msg);
}

function showTotalOrders(data){
    $("#result").html(data);
}

