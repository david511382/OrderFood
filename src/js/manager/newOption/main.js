var option = {Name:"新選單"};
init();

function init(){
    var newOptionIndex = -1;
    initShopName();
    InitCurrentOptionName(option.Name);
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
        Options:option.Name,
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
