function GetUserOrders(handler){
    $.ajax({
        type:"GET",
        url: "/api/order",  
    }).done(handler);
}

function ChangeView(shopID, handler){
    var data = {view:shopID};
    $.ajax({
        type:"PUT",
        url: "api/manager/changeshop",  
        data:data
    }).done(handler);
}

function GetShopMenu(shopID, handler){
    var url = 'api/manager/menu/shopmenu/' + shopID; // from treenode
    $.ajax({
        type:'GET',
        url: url
    }).done(handler);
}

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

function AddItem(shopID,name,price,handler){
    var url = 'api/manager/menu/item';
    var data = {
        shopID: shopID,
        name: name,
        price: price
    };
    $.ajax({
            type:'POST',
            url: url,
            data: data
        }).done(handler);
}

function AddOption(shopID,selectNum,items,selections,handler){
    var url = 'api/manager/menu/option';
    var data = {
        menuOptionJS:JSON.stringify({
        ShopID: shopID,
        SelectNum: selectNum,
        Items: items,
        Selections: selections
    })};
    $.ajax({
            type:'POST',
            url: url,
            data: data
        }).done(handler);
}

function DeleteOption(optionID, handler){
    var url =  "/api/manager/menu/option/" + optionID;
    $.ajax({
        type:"DELETE",
        url: url
    }).done(handler);
}

function AddSelection(optionID, name, price, handler){
    var url = 'api/manager/menu/selection';
    var data = {
        optionID: optionID,
        name: name,
        price: price
    };
    $.ajax({
            type:'POST',
            url: url,
            data: data
        }).done(handler);
}