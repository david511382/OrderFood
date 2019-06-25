init();

function init(){
    var newOptionIndex = -1;
    initShopName();
    InitCurrentOptionName("新選單");
    var select = CreateOptionNumSelect();
    InitItemTable();
    InitSelectionTable(newOptionIndex,);
}

function initShopName(){
    document.getElementById('shopNameInput').innerHTML = menuData.Shop.Name;   
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
