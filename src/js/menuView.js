var menuData;
var menu;

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

var names=[];

function init(kindData) {
    var fs=[];

    for (let i = 0; i < menuData.length; i++) {
        let kindData = menuData[i];
        makeKind(kindData,fs);
    }

    for (let i = 0; i < names.length; i++) {
        let moreButton = document.getElementById("more" + names[i])
        moreButton.addEventListener('click', fs[i]);
        }
    
}

function makeKind(kindData,fs){
    let items = kindData.Items;
    let sizes = kindData.Size;
    //let requiredSelection = kindData.RequiredSelection ;
    let checkOptions = kindData.CheckOption;
    for(let itemIndex=0;itemIndex<items.length;itemIndex++){
        let item= items[itemIndex];
        makeItem(item,checkOptions,fs,sizes);
    }
}

function getSelectOption(ID,sizes,prices){
    var select = document.createElement('select');
    select.id = ID;

    for(let i=0;i < sizes.length;i++){
        select.innerHTML+=getSelectOptionStr(sizes[i],prices[i]);
    }

    return select;
}

function getSelectOptionStr(text,value){
    return `<option value ="`+value+`">`+text+`</option>`;
}

function getCheckOption(name,options){
    var checkOption = document.createElement('span');
    checkOption.id = name+"CheckOption";

    for(let i=0;i < options.length;i++){
        let optionID = name+i.toString(10);
        checkOption.append(getCheckOptionStr(optionID, options[i].Name,options[i].Price))
    }
    return checkOption;
}

function getCheckOptionStr(ID,text,value){
    var checkSpan= document.createElement("span");
    
    var checkbox= document.createElement("input");
    checkbox.type = "checkbox";
    checkbox.id = ID;
    checkbox.value = value;
    checkSpan.appendChild(checkbox);

    checkSpan.innerHTML += text;
    return checkSpan;
}

function makeItem(item,checkOptionDatas,fs,sizes){
    var name  = item.Name;
    names.push(name);
    var prices     = item.Prices;
    var priceStrs=[];
    prices.forEach(element => {
        priceStrs.push(element.toString(10));
    });
    
    let rootDiv = document.createElement('div');
    let itemDiv = document.createElement('div');
    itemDiv.id = name;
    let itemInfoDiv = document.createElement('div');
    menu.appendChild(rootDiv);
    rootDiv.appendChild(itemDiv);
    itemDiv.appendChild(itemInfoDiv);

    var moreButton = document.createElement('button');
    moreButton.id = "more" + name;
    moreButton.type = "button";
    moreButton.classList.add("button");
    
    var span = document.createElement('span');
    span.innerText="More";
    moreButton.appendChild(span);
    
    fs.push(function() {
        more(itemDiv, function(name) {
            var namePrice = name + "Price";
            var nameSelect = name + "Select";
            var nameAmount = name + "Amount";

            var div = document.createElement('div');
            div.id = name;
            div.classList.add("order");
            itemDiv.appendChild(div);

            var removeButton = document.createElement('button');
            removeButton.type = "button";
            removeButton.classList.add("button");
            removeButton.addEventListener('click',function(){
                cancel(div);
            });
            removeButton.innerHTML= `<span>Remove</span>`;
            div.appendChild(removeButton);

            var priceSpan = document.createElement('span');
            priceSpan.innerHTML=" 價格: ";
            div.appendChild(priceSpan);
            
            var priceFont = document.createElement('font');
            priceFont.id=namePrice;
            priceSpan.appendChild(priceFont);

            var getSelectPrice = function(){
                if (sizes.length!==0){
                    return parseInt(selectOption.options[selectOption.selectedIndex].value);
                } else{
                    return prices[0];
                }
            }    
            var selectOption;       
            if (sizes.length!==0){
                selectOption = getSelectOption(nameSelect,sizes,prices);
                div.appendChild(selectOption);
            }
            var getSelectText = function(){
                if (sizes.length!==0){
                    return selectOption.options[selectOption.selectedIndex].innerHTML;
                } else{
                    return "";
                }
            }
            var selectObj={getText:getSelectText,getPrice:getSelectPrice};

            var checkOptionSpans = getCheckOption(name,checkOptionDatas);
            div.appendChild(checkOptionSpans);
            var checkOptions=[];
            for(let i=0;i < checkOptionSpans.childNodes.length;i++){
                checkOptions.push(checkOptionSpans.childNodes[i].childNodes[0]);
            }

            var amountSpan = document.createElement('span');
            amountSpan.innerHTML=" 數量";
            div.appendChild(amountSpan);

            var amountInput = document.createElement('input');
            amountSpan.appendChild(amountInput);
            amountInput.id = nameAmount;
            amountInput.type = "text";
            amountInput.value=1;
            amountInput.defaultValue = 1;
            amountInput.addEventListener('keypress', function(){
                if (window.event.keyCode==13){
                    onOrderChange(selectObj,checkOptions,amountInput,priceFont);
                    return false;
                };
             });            
            amountInput.onfocus=function(){this.select();};
            var text = document.createTextNode('個');
            amountSpan.appendChild(text);
            
            var f = function(v) {
                onOrderChange(selectObj, checkOptions, amountInput, priceFont);
            }

            if (sizes.length!==0){
                selectOption.selectedIndex++;
                selectOption.addEventListener('change', f);
            }
            checkOptions.forEach(element=>{
                element.addEventListener('change', f);
            });
            amountInput.addEventListener('change', f);

            calPrice(getSelectPrice, checkOptions, amountInput, priceFont);

            var msg = getText(getSelectText, checkOptions, amountInput, priceFont);
            addCart(name, msg);
        }
    );
    });

    

    itemInfoDiv.innerHTML +=getItemText(itemDiv.id,sizes,priceStrs);
    itemInfoDiv.appendChild(moreButton) + "</br > ";

    moreButton.addEventListener('click', fs);
}

function getItemText(name,sizes,priceStrs){
    let nameLen = name.length * 4;
    let spaceNum = maxSpace - nameLen;
    let spaces = "";
    for (; spaceNum > 0; spaceNum--) {
        spaces += spaceStr;
    }

    var text= name + spaces ;
    var hasSize=(sizes.length!==0)?true:false;

    for (let i=0;i < priceStrs.length;i++){
        if (hasSize)
            text+=sizes[i];
        text+=priceStrs[i]+ " ";
    }

    return text;
}