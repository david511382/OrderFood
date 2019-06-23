var menuData;
init();

function init(){
    var url = 'api/menu/shopmenu/' + ResopnseData; // from treenode
    $.ajax({
        type:'GET',
        url: url
    }).done(function(data){
        menuData = data;

        initShopName();
        initOptionButton();
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
    menuData.Options.forEach(function(menuOption) {
        var newTd = document.createElement('td');
        optionTableTr.appendChild(newTd);
        
        // add button
        var newButton = document.createElement('button');
        
        var id= 'none';
        if (menuOption.Option){
            id = menuOption.Option.ID;      
            
            // add remove option button to td
            var newRmButton = document.createElement('button');
            newRmButton.innerHTML = "-";
            newTd.appendChild(newRmButton)
        }
        newButton.Name = id + "OptionButton";
        newButton.innerHTML = menuOption.Name;
        optionButtonDiv.appendChild(newButton);
      }); 

    
    var tdForAddOption = document.createElement('td');
    optionTableTr.appendChild(tdForAddOption);
}

function newOptionButtonClick(){
    // 換頁新增
    // var url = 'api/menu/option';
    // var data = {
    //     selectionName: "預設選項",
    //     selectNum: 0
    // };
    // $.ajax({
    //         type:'POST',
    //         url: url,
    //         data: data
    //     }).done(function(option){
    //         alert(option);
    //     });
}

function newItem(){
    var itemNameInput = document.getElementById('newItemNameInput');
    var itemPriceInput = document.getElementById('newItemPriceInput');
    var name = itemPriceInput.value;
    var price = itemPriceInput.value;

    if (!name){
        alert("please input name!");
        return ;
    }
    if (!price){
        price = 0;
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
        });
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

