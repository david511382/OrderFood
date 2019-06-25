// selectedOptionIndex -1 : new option mode
// selectedOptionIndex 0 : all option mode
// selectedOptionIndex 1+ : normal option mode

var menuData;
var selectedOptionIndex;
init();

function init(){
    GetShopMenu(ResopnseData,function(data){
        menuData = data;
        selectedOptionIndex = 0;
        initShopName();
        InitCurrentOptionName();
        initOptionButton();
        initOptionNumSelect();
        menuOption = menuData.Options[selectedOptionIndex]
        InitItemTable(menuOption.Items);
        InitSelectionTable(selectedOptionIndex,menuOption.Selections);
    });
}

function initShopName(){
    document.getElementById('shopNameInput').value = menuData.Shop.Name;   
}

function initOptionButton(){
    var optionButtonDiv = document.getElementById('optionButtonDiv');
    var optionTableTr = document.getElementById('optionTableTr');    
    
    if (!menuData.Options){
        return 
    }
    for (let index = 0; index < menuData.Options.length;index++){
        let menuOption = menuData.Options[index];
        var newTd = document.createElement('td');
        optionTableTr.appendChild(newTd);
        
        // add button
        var newButton = document.createElement('button');
        
        var id= 'none';
        if (menuOption.Option){            
            // add remove option button to td
            var newRmButton = document.createElement('button');
            newRmButton.innerHTML = "-";
            newTd.appendChild(newRmButton)

            id = index;
        }
        newButton.Name = id + "OptionButton";
        newButton.innerHTML = menuOption.Name;
        newButton.addEventListener('click',function(){
            optionButtonClick(index);
        });

        optionButtonDiv.appendChild(newButton);
    }
    
    var tdForAddOption = document.createElement('td');
    optionTableTr.appendChild(tdForAddOption);
}

function optionButtonClick(index){
    selectedOptionIndex = index
    InitCurrentOptionName();
    initOptionNumSelect();
    menuOption = menuData.Options[selectedOptionIndex]
    InitItemTable(menuOption.Items);
    InitSelectionTable(selectedOptionIndex, menuOption.Selections);
}

function CreateOptionNumSelect(){
    var optionSelectTd = document.getElementById('optionSelectTd');
    var a = document.createElement('a');
    a.innerHTML = "必選數量";
    optionSelectTd.appendChild(a);


    var select = document.createElement('select');
    select.innerHTML = "必選數量";
    optionSelectTd.appendChild(select);

    var option = document.createElement('option');
    option.value = 0;
    option.innerHTML = "0";
    select.options.add(option);    

    return select;
}

function initOptionNumSelect(){
    // clear
    var optionSelectTd = document.getElementById('optionSelectTd');
    optionSelectTd.innerHTML = "";
            
    menuOption = menuData.Options[selectedOptionIndex];
    menuSelections = menuOption.Selections;
    if (!menuSelections){
        return ;
    }
    
    var select = CreateOptionNumSelect();
    for (let i = 1; i <= menuSelections.length; i++){
        let option = document.createElement('option');
        option.value = i;
        option.innerHTML = i;
        select.options.add(option);
    }    
}

function InitItemTable(items,addItemButtonClick){    
    var addItemButton = document.getElementById('addItemButton');
    if (addItemButtonClick===undefined){
        addItemButton.onclick="newItemButtonClick()";
    }else{
        addItemButton.onclick=addItemButtonClick;
    }

    var itemTable = document.getElementById('itemTable');
    
    // clear 
    for (;2< itemTable.childNodes.length;){
        itemTable.removeChild(itemTable.lastChild);
    }

    if (!items){
        return;
    }
    items.forEach(function(item) {
        var newTr = NewItemTableTr(item,null);
        itemTable.appendChild(newTr);
    }); 
}

function NewItemTableTr(item, deleteButtonClick){
    var newTr = document.createElement('tr');

    var newTd = document.createElement('td');
    newTd.innerHTML = item.Options;
    newTr.appendChild(newTd);

    newTd = document.createElement('td');
    newTd.innerHTML = item.Name;
    newTr.appendChild(newTd);

    newTd = document.createElement('td');
    newTd.innerHTML = item.Price;
    newTr.appendChild(newTd);

    newTd = document.createElement('td');
    var newButton = document.createElement('button');
    newButton.innerHTML ="刪除";
    newButton.onclick = deleteButtonClick;
    newTr.appendChild(newButton);

    return newTr;
}

function InitCurrentOptionName(name){
    var currentOptionNameA = document.getElementById('currentOptionNameA');
    var itemOptionNameTd = document.getElementById('newItemOptionNameTd');
    
    var itemOptionName = name;
    if (name === undefined){
        menuOption = menuData.Options[selectedOptionIndex];
        name = menuOption.Name;

        if (menuOption.Option){
            itemOptionName = name;
        }else{
            itemOptionName = "";
        }
    }

    currentOptionNameA.innerHTML = name;
    itemOptionNameTd.innerHTML = itemOptionName;
}
function newItemButtonClick(){
    var itemNameInput = document.getElementById('newItemNameInput');
    var name = itemNameInput.value;
    if (!name){
        alert("please input name!");
        return ;
    }

    var itemPriceInput = document.getElementById('newItemPriceInput');
    var price =parseInt(itemPriceInput.value);
    if (isNaN(price)){
        alert("please input integer price");
        return;
    }


    AddItem(menuData.Shop.ID,name,price,function(data){
        init();
    });
}

function InitSelectionTable(selectedOptionIndex,selections, newSelectionButtonClick){
    var selectionTableTd = document.getElementById('selectionTableTd');
    // clear 
    selectionTableTd.innerHTML = "";

    if (selectedOptionIndex === 0){
        return ;
    }

    // init
    var a = document.createElement('a');
    a.innerHTML = "選單選項";
    selectionTableTd.appendChild(a);
    var selectionTable = document.createElement('table');
    selectionTable.id = "selectionTable"
    selectionTableTd.appendChild(selectionTable);

    // first row
    var newTr = document.createElement('tr');
    selectionTable.appendChild(newTr);
    
    var newTd = document.createElement('td');
    newTd.innerHTML = "名稱";
    newTr.appendChild(newTd);

    newTd = document.createElement('td');
    newTd.innerHTML = "加價";
    newTr.appendChild(newTd);

    newTd = document.createElement('td');
    newTd.innerHTML = "操作";
    newTr.appendChild(newTd);

    // second row
    newTr = document.createElement('tr');
    selectionTable.appendChild(newTr);
    
    newTd = document.createElement('td');
    var newInput = document.createElement('input');
    newInput.id = "newSelectionNameInput";
    newTd.appendChild(newInput);
    newTr.appendChild(newTd);

    newTd = document.createElement('td');
    newInput = document.createElement('input');
    newInput.id = "newSelectionPriceInput";
    newTd.appendChild(newInput);
    newTr.appendChild(newTd);

    newTd = document.createElement('td');
    var newButton = document.createElement('button');
    newButton.id = "addSelectionButton";
    newButton.innerHTML="新增";
    if (newSelectionButtonClick === undefined){
        newButton.onclick =this.newSelectionButtonClick;    
    }else{
        newButton.onclick =newSelectionButtonClick;
    }
    newTd.appendChild(newButton);
    newTr.appendChild(newTd);


    if (!selections){
        return;
    }
    selections.forEach(function(selection) {
        selection.Price = (selection.Price)? selection.Price : 0;
        var newTr =NewSelectionTableTr(selection)
        selectionTable.appendChild(newTr);
        }); 
}


function NewSelectionTableTr(selection, deleteButtonClick){
    var newTr = document.createElement('tr');

    var newTd = document.createElement('td');
    newTd.innerHTML = selection.Name;
    newTr.appendChild(newTd);

    newTd = document.createElement('td');
    var price = (selection.Price)? selection.Price : 0;
    newTd.innerHTML = price;
    newTr.appendChild(newTd);

    var newButton = document.createElement('button');
    newButton.innerHTML ="刪除";
    if (deleteButtonClick === undefined){
        newButton.onclick = this.deleteButtonClick;    
    }else{
        newButton.onclick = deleteButtonClick;
    }
    newTr.appendChild(newButton);

    return newTr;
}

function newSelectionButtonClick(){
    var selectionNameInput = document.getElementById('newSelectionNameInput');
    var name = selectionNameInput.value;
    if (!name){
        alert("please input name!");
        return ;
    }

    var selectionPriceInput = document.getElementById('newSelectionPriceInput');
    var price =parseInt(selectionPriceInput.value);
    if (isNaN(price)){
        alert("please input integer price");
        return;
    }

    AddSelection(menuData.Shop.ID,name,price,function(data){
        init();
    });
}

function newOptionButtonClick(){
    var url = 'manager/newoption';
    $.ajax({
            type:'GET',
            url: url
        }).done(UpdatePage);
}

function shopNameInputKeyPress(shopName){
    if (window.event.keyCode==13){
        shopID = menuData.Shop.ID;        
        if (!shopID){
            alert("delete err");
            return false;
        }      

        oldShopName = menuData.Shop.Name;     
        if (!shopName){
            alert("err")
            return false;
        }else if (shopName == ""){
            alert("please input shop name!")
            return false;
        }else if (shopName == oldShopName){
            return false
        }

        UpdateShop(shopID,shopName,
            function(success){
                if (!success){
                    alert('fail');
                    return ;
                }
    
                var url = '/manager/menutree';
                $.ajax({
                    type:'GET',
                    url: url
                }).done(UpdatePage);
    
                alert('修改商店成功!');
            }
        )

        return false;
    };
}

function removeShopButtonClick(){
    shopID = menuData.Shop.ID;        
    if (!shopID){
        alert("delete err");
        return 
    }

    DeleteShop(
        shopID,// from tree node js
        function(success){
            if (!success){
                alert('fail');
                return ;
            }

            var url = '/manager/menutree';
            $.ajax({
                type:'GET',
                url: url
            }).done(UpdatePage);

            toNewShop();
            
            alert('刪除商店成功!');
        }	
    );
}

