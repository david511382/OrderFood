var id = 0;
var shopcart = document.getElementById("shopcart");
shopcart.value = "";
var shopcartName = "cart"

var shopTotalPrice = document.getElementById("total");
shopTotalPrice.innerHTML = 0;
var total = 0;

function more(item, addItems) {
    id++;
    var idName = "order" + id.toString(10);

    addItems(idName);

    if (item.childNodes.length === 2)
        item.parentNode.appendChild(document.createElement('p'));
}

function addCart(name, msg) {
    var div = document.createElement('div');
    div.id = name + shopcartName;
    div.innerHTML = msg

    shopcart.appendChild(div);
}

function changeCart(name, msg) {
    var div = document.getElementById(name + shopcartName);
    div.innerHTML = msg;
}

function cancel(orderDiv) {
    var orders = orderDiv.parentNode;
    if (orders.childNodes.length === 2) {
        let blank = orders.parentNode.childNodes[1];
        blank.parentNode.removeChild(blank);
    }


    var namePrice = orderDiv.id + "Price";
    var price = document.getElementById(namePrice);
    total -= parseInt(price.innerHTML);
    shopTotalPrice.innerHTML = total;

    var shopcartOrderName = orderDiv.id + shopcartName;

    orderDiv.parentNode.removeChild(orderDiv);

    var shopcartOrder = document.getElementById(shopcartOrderName);
    shopcartOrder.parentNode.removeChild(shopcartOrder);
}

function onOrderChange(selectObj, options, amount, price) {
    var amountValue = parseInt(amount.value);
    if (amountValue<1 || isNaN(amountValue)){
        alert("數量不得<1");
        amount.value=1;
        return
    }else{
        amount.value=amountValue;
    }

    calPrice(selectObj.getPrice, options, amount, price);

    var msg = getText(selectObj.getText, options, amount, price);
    changeCart(price.parentNode.parentNode.id, msg);
}

function calPrice(getSelectPrice, options, amount, price) {
    var totalPrice=getSelectPrice();

    for(let i=0;i<options.length;i++){
        let option = options[i];
        if (option.checked)
            totalPrice += parseInt(option.value);
    }
    
    totalPrice *= parseInt(amount.value);

    var oldCost = price.innerHTML;
    price.innerHTML = totalPrice;

    total += totalPrice - oldCost;
    shopTotalPrice.innerHTML = total;
}


function getText(selectText,options,amount,price){
    var name = price.parentNode.parentNode.parentNode.id;
    var quen=amount.value;
    var cost=parseInt(price.innerHTML);
    
    var withStr="";
    for(let i=0;i<options.length;i++){
        let option = options[i];
        if (option.checked)
            withStr += option.parentNode.childNodes[1].data + " ";
    }

    return name+" "+selectText()+" "+withStr +" "+quen+" 個 "+ price.innerHTML + `元\n`;
}

