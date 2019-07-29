import Vuex from 'vuex'

const appStore = () => {
    return new Vuex.Store({
        state: {
            menu_list: {}
        },
        mutations: {
            menu_list_update(state, payload){
                state/menu_list = {...payload.data}
            }
        }
    })
}

export default appStore
