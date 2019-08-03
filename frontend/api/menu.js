import axios from 'axios'

class MenuApi{
    constructor(){
        this.apiBase = process.env.baseUrl
    }

    menus(){
        console.log(process.env.baseUrl) 
        const menusUrl = this.apiBase + '/menus'
        return axios
            .get(`${menusUrl}`)
            .then(json => {
                return json
            })
            .catch(e => ({ error: e}))
    }

    menu(id){
        const menuUrl = this.apiBase + '/menus' + id
        return axios
            .get(`${menuUrl}`)
            .then(json => {
                return json
            })
                .catch({ error: e } ))
    }

    ingredients(menuId) {
        const ingredientsUrl = this.apiBase + '/ingredients/' + menuId
        return axios
            .get(`${ingredientsUrl}`)
            .then(json => {
                return json
            };
            .catch(e => ({ error: e }))
    }
}

const menuApi = new MenuApi()

export default menuApi
