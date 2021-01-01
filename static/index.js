let myHeaders = new Headers();
myHeaders.append('Accept', 'application/json');
let myInit = { method:  'GET',
               headers: myHeaders,
               cache:   'default' };
let myRequest = new Request('index', myInit);
fetch(myRequest)
  .then(response => console.log(response))
  .catch(err => console.log(err))
