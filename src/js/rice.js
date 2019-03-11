var menuData;
var menu;

var script = document.createElement("SCRIPT");
$.ajax({
        type: "POST",
        url: "/get/menu"
    }).done(function(data) {
        menuData = data;
        menu = document.getElementById("menu");
        init();
    });

const menuId = "menu";
const maxSpace = 30;
const spaceStr = `&nbsp;`
const plusPrice = 10;


function init() {
    var fs=[];

    for (let i = 0; i < menuData.length; i++) {
        let itemData = menuData[i];
        let rootDiv = document.createElement('div');
        let itemDiv = document.createElement('div');
        itemDiv.id = itemData.Name;
        let itemInfoDiv = document.createElement('div');
        menu.appendChild(rootDiv);
        rootDiv.appendChild(itemDiv);
        itemDiv.appendChild(itemInfoDiv);

        var moreButton = document.createElement('button');
        itemInfoDiv.appendChild(moreButton);
        moreButton.type = "button";
        moreButton.innerText = "more";
        moreButton.id = "more" + i.toString(10);
        
        let halfPriceStr = itemData.Half.toString(10);
        let normalPriceStr = itemData.Normal.toString(10);
        let largePriceStr = (itemData.Normal + plusPrice).toString(10);

fs.push(function() {
    more(itemDiv, function(name) {
        var nameRice = name + "Rice";
        var nameCheese = name + "Cheese";
        var namePrice = name + "Price";
        var nameSelect = name + "Select";
        var nameAmount = name + "Amount";

        var div = document.createElement('div');
        div.id = name;
        div.innerHTML = ` <font>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</font>
        <button type="button" onclick="cancel(`+name+`)">delete</button > 價格: <font id = "`+namePrice+`" > </font> 
    <select id="`+nameSelect+`">
        <option value ="`+halfPriceStr+`">半捲</option> <option value = "`+normalPriceStr+`"> 正常 </option>
        <option value ="`+largePriceStr+`">加大</option></select>
    <input type="checkbox" id="`+nameRice+`" name="`+nameRice+`" value="5">換拾穀米5  
    <input type="checkbox" id="`+nameCheese+`" name="`+nameCheese+`" value="10">加起司10  
    數量<input id="`+nameAmount+`" type="text" name="ammount" value="1" onfocus="this.select()" onkeypress="onOrderChang(`+nameSelect+`,`+nameRice+`,`+nameCheese+`,`+nameAmount+`,`+namePrice+`); if (window.event.keyCode==13) return false;">個`
        
             itemDiv.appendChild(div);
             
                var select = document.getElementById(nameSelect);
                var rice = document.getElementById(nameRice);
                var cheese = document.getElementById(nameCheese);
                var amount = document.getElementById(nameAmount);
                var price = document.getElementById(namePrice);

                select.selectedIndex++;
                    var f = function() {
                        onOrderChang(select, rice, cheese, amount, price);
                    }
                    select.addEventListener('change', f);
                    rice.addEventListener('change', f);
                    cheese.addEventListener('change', f);
                    amount.addEventListener('change', f);
                
                    orderPrice(select, rice, cheese, amount, price);

                var msg = makeMsg(select, rice, cheese, amount, price);
                addCart(name, msg);
      });
    });

        let nameLen = itemDiv.id.length * 4;
        let spaceNum = maxSpace - nameLen;
        let spaces = "";
        for (; spaceNum > 0; spaceNum--) {
            spaces += spaceStr;
        }

        itemInfoDiv.innerHTML += itemDiv.id + spaces + "半捲" + halfPriceStr + " 正常" + normalPriceStr + " 加大" + largePriceStr + "</br > ";

    }


    for (let i = 0; i < menuData.length; i++) {
        let moreButton = document.getElementById("more" + i.toString(10))
        moreButton.addEventListener('click', fs[i]);
        }
    
}