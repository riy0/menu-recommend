import axios from 'axios'

class MenuApi{
    constructor(){
        this.apiBase = 'http://127.0.0.1:8080/menus'
    }

    menus(){
        return axios
        .get(`${this.apiBase}`)
        .then(json => {
            return json
        })
        .catch(e => ({ error: e}))
    }
}

const menuApi = new MenuApi()

export default menuApi
