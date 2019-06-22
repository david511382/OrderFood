function NewShopButtonClick(){
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

