function AddShop(shopName, handler){
    if (!shopName || shopName === ""){
        alert("please input shop name!");
        return ;
    }

    var data = {name:shopName};
    $.ajax({
        type: "POST",
        url: "/api/menu/shop",  
        data: data
    }).done(handler);
}

