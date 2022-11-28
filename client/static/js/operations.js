function request(url){
    return new Promise((resolve, reject)=>{
       $.ajax(url,{
            dataType: 'json',
            timeout: 30000,
            success: (data, status, xhr) =>{
                resolve(data.data);
            },
            error: (jqXhr)=>{
                showAlertModal('Error', jqXhr.responseJSON.message);
                reject(jqXhr.responseJSON)
            }
       })
    });
}

function request(url, method, data){
    return new Promise((resolve, reject)=>{
       $.ajax(url,{
            method: method,
            data: JSON.stringify(data),
            dataType: 'json',
            timeout: 30000,
            success: (data, status, xhr) =>{
                resolve(data.data);
            },
            error: (jqXhr)=>{
                showAlertModal('Error', jqXhr.responseJSON.message);
                reject(jqXhr.responseJSON)
            }
       })
    });
}

function invalidDate(date){
    let formatDate = new Date(date);
    if(formatDate.getFullYear() == 0){
        return '';
    }
    
    return formatDate.toLocaleDateString();
}

function showAlertModal(title, message){
    $('#alertModal').find('.modal-title').html(title);
    $('#alertModal').find('.modal-body').find('article').html(message);
    $('#alertModal').modal('show');
}

function confirmModal(title, message, yesCallback, noCallback){
    $('#confirmModal').find('.modal-title').html(title);
    $('#confirmModal').find('.modal-body').find('article').html(message);
    $('#confirmModal').find('.noButton').click(noCallback);
    $('#confirmModal').find('.yesButton').click(yesCallback);
    $('#confirmModal').modal('show');
}