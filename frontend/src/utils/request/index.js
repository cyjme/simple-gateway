import axios from 'axios';

let request = axios;
request.defaults.baseURL="http://localhost:11111";

export default request;
