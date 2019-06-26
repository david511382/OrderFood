const optionName = "新選單";
init();

function init(){
    var newOptionIndex = -1;
    initShopName();
    InitCurrentOptionName(optionName);
    var select = CreateOptionNumSelect();
    InitItemTable(null,newItemButtonClick);
    InitSelectionTable(newOptionIndex,null,newSelectionButtonClick);
}

function initShopName(){
    document.getElementById('shopNameInput').innerHTML = menuData.Shop.Name;   
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


    var itemTable = document.getElementById('itemTable');
    
    var item = {
        Options:optionName,
        Name:name,
        Price:price
    };
    var newTr = NewItemTableTr(item)
    itemTable.appendChild(newTr);
    
    itemNameInput.value = "";
    itemPriceInput.value = "";
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

    var selection = {
        Name:name,
        Price:price
    };
    var selectionTable = document.getElementById('selectionTable');
    var newTr =NewSelectionTableTr(selection)
    selectionTable.appendChild(newTr);
    
    selectionNameInput.value = "";
    selectionPriceInput.value = "";
}

function doneButtonClick(){
    var shopID = menuData.Shop.ID

    var selectNumSelect = document.getElementById("selectNumSelect")

    var itemTable = document.getElementById("itemTable")
    var items = [];
    for (i = 2;i < itemTable.childNodes.length;i++){
        var tr = itemTable.childNodes[i];
        var name = tr.childNodes[1].innerHTML;
        var price = parseInt(tr.childNodes[2].innerHTML);
        items.push({name:name,price:price});
    }
    if (items.length === 0){
        alert("at least one item");
        return ;
    }

    var selectionTable = document.getElementById("selectionTable")
    var selections = [];
    for (i = 2;i < selectionTable.childNodes.length;i++){
        var tr = selectionTable.childNodes[i];
        var name = tr.childNodes[0].innerHTML;
        var price = parseInt(tr.childNodes[1].innerHTML);
        selections.push({name:name,price:price});
    }
    if (selections.length === 0){
        alert("at least one selection");
        return ;
    }

    AddOption(
        shopID,
        selectNumSelect.value,
        items,
        selections,
        function(menu){
            if (!menu){
                alert('fail');
                return ;
            }

            var url = '/manager/menutree';
            $.ajax({
                type:'GET',
                url: url
            }).done(UpdatePage);

            alert('成功!');

            //menuData.Options.push(menu);
            
            // var newOptionIndex = menuData.Options.length - 1;
            // LoadManageMenu(menuData, newOptionIndex);
            toManageShop(shopID);
        }	
    );
}

function cancelButtonClick(){
    AddShop(
        document.getElementById("shopNameInput").value,
        function(result){
            if (!result){
                alert('fail');
                return ;
            }

            var url = '/manager/menutree';

            $.ajax({
                type:'GET',
                url: url
            }).done(UpdatePage);

            alert('新增商店 ' + result.Name + ' 成功!');
            toManageShop(result.ID);
        }	
    );
}
