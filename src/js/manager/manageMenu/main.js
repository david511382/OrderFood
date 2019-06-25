var menuData;
var selectedOptionIndex;
init();

function init(){
    var url = 'api/menu/shopmenu/' + ResopnseData; // from treenode
    $.ajax({
        type:'GET',
        url: url
    }).done(function(data){
        menuData = data;
        selectedOptionIndex = 0;
        InitShopName();
        initCurrentOptionName();
        initOptionButton();
        initOptionNumSelect();
        initItemTable();
        initSelectionTable();
    });
}

function InitShopName(){
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
    initCurrentOptionName();
    initOptionNumSelect();
    initItemTable();
    initSelectionTable();
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

function initItemTable(){    
    var itemTable = document.getElementById('itemTable');
    
    // clear 
    for (;2< itemTable.childNodes.length;){
        itemTable.removeChild(itemTable.lastChild);
    }

    menuOption = menuData.Options[selectedOptionIndex];
    menuOption.Items.forEach(function(item) {
        var newTr = document.createElement('tr');
        itemTable.appendChild(newTr);

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
        newButton.onclick ="function t(" + item.ID + ")";
        newTr.appendChild(newButton);
      }); 
}

function initCurrentOptionName(){
    var currentOptionNameA = document.getElementById('currentOptionNameA');
    var itemOptionNameTd = document.getElementById('newItemOptionNameTd');
    
    menuOption = menuData.Options[selectedOptionIndex];
    var name = menuOption.Name;

    currentOptionNameA.innerHTML = name;
    
    if (menuOption.Option){
        itemOptionNameTd.innerHTML = name;
    }else{
        itemOptionNameTd.innerHTML = "";
    }
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


    var url = 'api/menu/item';
    var data = {
        shopID: menuData.Shop.ID,
        name: name,
        price: price
    };
    $.ajax({
            type:'POST',
            url: url,
            data: data
        }).done(function(data){
            init();
        });
}

function initSelectionTable(){
    var selectionTable = document.getElementById('selectionTable');
    
    // clear 
    for (;2< selectionTable.childNodes.length;){
        selectionTable.removeChild(selectionTable.lastChild);
    }

    menuOption = menuData.Options[selectedOptionIndex];
    if (!menuOption.Selections){
        return;
    }
    menuOption.Selections.forEach(function(selection) {
        var newTr = document.createElement('tr');
        selectionTable.appendChild(newTr);

        var newTd = document.createElement('td');
        newTd.innerHTML = selection.Name;
        newTr.appendChild(newTd);

        newTd = document.createElement('td');
        var price = (selection.Price)? selection.Price : 0;
        newTd.innerHTML = price;
        newTr.appendChild(newTd);

        newTd = document.createElement('td');
        var newButton = document.createElement('button');
        newButton.innerHTML ="刪除";
        newButton.onclick ="function t(" + selection.ID + ")";
        newTr.appendChild(newButton);
      }); 
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

    var url = 'api/menu/selection';
    var data = {
        optionID: menuData.Shop.ID,
        name: name,
        price: price
    };
    $.ajax({
            type:'POST',
            url: url,
            data: data
        }).done(function(data){
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

