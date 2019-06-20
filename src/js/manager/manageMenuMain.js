var shopData;
//var selectedShopID = %d;

$.ajax({
    type:"GET",
    url: "/menu/shopmenu"
}).done(init);

function init(data){
    shopData = data
}