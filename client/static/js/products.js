$(document).ready(()=>{
    loadData();
    $('#btnCreate').click(()=>{
        showModal(1, 0);
    });

    $('#btnEdit').click(()=>{
        let tr = $('tr.selected');
        if(tr.length > 0){   
            showModal(2, parseInt(tr.data('id')));
        }else{
            alert('Select a row to update');
        }
    });

    $('#productsTable').on('click','td', (e)=>{
        let td = e.currentTarget;
        let tr = $(td).closest('tr');

        if(tr.hasClass('selected')){
            $('tr.selected').removeClass('selected');
        }else{
            $('tr.selected').removeClass('selected');
            $(tr).addClass('selected');
        }
    });

    $('#btnSave').click(()=>{
        let type = $('#modalProducts').data('type');
        let data = {
            Id: $('#txId').val(),
            Name: $('#txName').val(),
            UnitPrice: parseFloat($('#txUnitPrice').val()),
            Tax: parseFloat($('#txTax').val())
        };
        
        let url = 'api/products';
        let method = 'POST';
        if(type == 2){
            url += '/' + data.Id;
            method = 'PUT';
        }

        request(url, method, data).then(res=>{
            showAlertModal('Information', 'Changes saved succesfully');
            $('#modalProducts').modal('hide');
            loadData();
        });
    });

    $('#btnDelete').click(()=>{
        let tr = $('tr.selected');
        if(tr.length > 0){   
            confirmModal('Delete Customer', 'Are you sure you want to delete this record?', ()=>{
                request('api/products/' + tr.data('id'), 'DELETE', null).then(res=>{
                    showAlertModal('Information', 'Record removed succesfully');
                    loadData();
                });
            }, ()=>{});
        }else{
            showAlertModal('Warning', 'Select a row to update');
        }
    });

    function loadData(){
        request('api/products').then(res=>{
            let tBody = $('#productsTable').find('tbody');
            tBody.html('');
            for(let item of res){
                tBody.append('<tr data-id="' + item.Id + '"><td>' + item.Id + '</td>' + '<td>' + item.Name + '</td>' + '<td>' + item.UnitPrice + '</td>' + '<td>' + item.Tax + '</td>' + '<td>' + new Date(item.CreationDate).toLocaleDateString() + '</td>' + '<td>' + invalidDate(item.ModificationDate) + '</td></tr>');
            }
        }).catch(err=>{
            alert(err);
        });
    }

    function showModal(type, id){
        let title = type == 1 ? 'Create Customer' : 'Edit Customer';
        $('#modalProducts').data('type', type);
        $('#modalProducts').find('.modal-title').html(title);
        if(type == 2){
            request('api/products/' + id).then(res=>{
                $('#txId').val(res.Id);
                $('#txName').val(res.Name);
                $('#txUnitPrice').val(res.UnitPrice);
                $('#txTax').val(res.Tax);
                $('#modalProducts').modal('show');
            });
        }else{
            $('#txId').val('');
            $('#txName').val('');
            $('#txUnitPrice').val('');
            $('#txTax').val('');
            $('#modalProducts').modal('show');
        }
    }
});