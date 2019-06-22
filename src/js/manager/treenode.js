function toHome(){
    $.ajax({
        type:"GET",
        url: "/manager"
    }).done(changeWholePage);
}

function toNewShop(){
    var url =  "/manager/newshop";
    
    $.ajax({
        type:"GET",
        url: url
    }).done(UpdatePage);
}

function toManageShop(shopID){
    var url =  "/manager/managemenu?shopID=";
    if (shopID !== undefined) {
        url += shopID;
    }
    
    $.ajax({
        type:"GET",
        url: url
    }).done(UpdatePage);
}

function changeWholePage(html){
    document.body.innerHTML = html;
}