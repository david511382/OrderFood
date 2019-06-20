var shopData;
//var selectedShopID = %d;

$.ajax({
    type:"GET",
    url: "/api/menu/shopmenu"
}).done(init);

function init(data){
    shopData = data
}