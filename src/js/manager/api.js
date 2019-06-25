function AddShop(shopName, handler){
    if (!shopName || shopName === ""){
        alert("please input shop name!");
        return ;
    }

    var data = {name:shopName};
    $.ajax({
        type: "POST",
        url: "/api/manager/menu/shop",  
        data: data
    }).done(handler);
}

function UpdateShop(shopID, shopName, handler){
    var url =  "/api/manager/menu/shop/" + shopID;
    var data = {name:shopName}
    $.ajax({
        type:"PUT",
        url: url,
        data: data
    }).done(handler);
}

function DeleteShop(shopID, handler){
    var url =  "/api/manager/menu/shop/" + shopID;
    $.ajax({
        type:"DELETE",
        url: url
    }).done(handler);
}