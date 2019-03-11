var id = 0;
var shopcart = document.getElementById("shopcart");
shopcart.value = "";
var shopcartName = "cart"

var shopTotalPrice = document.getElementById("total");
shopTotalPrice.innerHTML = 0;
var total = 0;

// function more(item, half, normal, plus) {
//     id++;
//     var idName = "order" + id.toString(10);

//     getDiv(idName, half, normal, plus, function(div, nameSelect, nameRice, nameCheese, nameAmount, namePrice) {
//         item.appendChild(div);

//         var select = document.getElementById(nameSelect);
//         var rice = document.getElementById(nameRice);
//         var cheese = document.getElementById(nameCheese);
//         var amount = document.getElementById(nameAmount);
//         var price = document.getElementById(namePrice);

//         initDiv(select, rice, cheese, amount, price);
//         var msg = makeMsg(select, rice, cheese, amount, price);
//         addCart(idName, msg);
//     });

//     if (item.childNodes.length === 2)
//         item.parentNode.appendChild(document.createElement('p'));
// }
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

function orderPrice(select, rice, cheese, amount, price) {
    var totalPrice = parseInt(select.options[select.selectedIndex].value);

    if (rice.checked)
        totalPrice += parseInt(rice.value);

    if (cheese.checked)
        totalPrice += parseInt(cheese.value);

    totalPrice *= parseInt(amount.value);

    var oldCost = price.innerHTML;
    price.innerHTML = totalPrice;

    total += totalPrice - oldCost;
    shopTotalPrice.innerHTML = total;
}

// function initDiv(select, rice, cheese, amount, price) {
//     select.selectedIndex++;

//     var f = function() {
//         onOrderChang(select, rice, cheese, amount, price);
//     }
//     select.addEventListener('change', f);
//     rice.addEventListener('change', f);
//     cheese.addEventListener('change', f);
//     amount.addEventListener('change', f);

//     orderPrice(select, rice, cheese, amount, price);
// }

function onOrderChang(select, rice, cheese, amount, price) {
    orderPrice(select, rice, cheese, amount, price);

    var msg = makeMsg(select, rice, cheese, amount, price);
    changeCart(select.parentNode.id, msg);
}

// function getDiv(name, half, normal, plus, done) {
//     var nameRice = name + "Rice";
//     var nameCheese = name + "Cheese";
//     var namePrice = name + "Price";
//     var nameSelect = name + "Select";
//     var nameAmount = name + "Amount";

//     var div = document.createElement('div');
//     div.id = name;
//     div.innerHTML = ` < font > & nbsp; & nbsp; & nbsp; & nbsp; & nbsp; < /font>
//     <button type="button" onclick="cancel(`+name+`)">delete</button > 價格: < font id = "`+namePrice+`" > < /font> 
//     <select id="`+nameSelect+`">
//         <option value ="`+half+`">半捲</option > < option value = "`+normal+`" > 正常 < /option>
//         <option value ="`+plus+`">加大</option > < /select>
//     <input type="checkbox" id="`+nameRice+`" name="`+nameRice+`" value="5">換拾穀米5  
//     <input type="checkbox" id="`+nameCheese+`" name="`+nameCheese+`" value="10">加起司10  
//     數量<input id="`+nameAmount+`" type="text" name="ammount" value="1" onfocus="this.select()" onkeypress="onOrderChang(`+nameSelect+`,`+nameRice+`,`+nameCheese+`,`+nameAmount+`,`+namePrice+`); if (window.event.keyCode==13) return false;">個`
    
//     done(div,nameSelect,nameRice,nameCheese,nameAmount,namePrice);
// }

function makeMsg(select,rice,cheese,amount,price){
    var name = select.parentNode.parentNode.id;
    var size=select.options[select.selectedIndex].innerHTML;
    var quen=amount.value;
    var cost=parseInt(price.innerHTML);
    
    var withRice="";
    if (rice.checked)
        withRice="換拾穀米 ";
        
    var withCheese="";
    if (cheese.checked)
        withCheese="加起司 "

    return name+" "+size+" "+withRice+withCheese+quen+" 個 "+ price.innerHTML + `元\n`;
}