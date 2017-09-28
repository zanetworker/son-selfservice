import {UPDATE_MODAL, UPDATE_LOADING} from '../actions'

//e.g., state
// {
//     id: "12345",
//     name: "Firewall",
//     state: "stopped"
// }
const INIT_STATE = {
      modalStateBasic: false,
      loadingStateBasic: true,
      modalStateAnon: false,
      loadingStateAnon: true
};



export default (state = INIT_STATE, action) => {

  switch (action.type){
    case UPDATE_MODAL:
      let theType = action.payload.type;
      if (theType === "basic"){
        console.log("Updating Modal Basic" + action.payload.state)
        return createModalStateObjectBasic(action.payload.state)
      }


     if (theType === "anon"){
        console.log("Updating Modal Anon" + action.payload.state)
        return createModalStateObjectAnon(action.payload.state)
      }
      break;

    case UPDATE_LOADING:
        return createLoadingStateObjectBasic(action.payload)
        break;

    default:
      return state;
      }
  }


// Basic Service Type Object Creation
const createModalStateObjectBasic= (modalState) => {
  return {
      modalStateBasic: modalState,
      loadingStateBasic: INIT_STATE.loadingStateBasic
  }
}

  const createLoadingStateObjectBasic= (loadingState) => {
    return {
        modalStateBasic: INIT_STATE.modalStateBasic,
        loadingStateBasic: loadingState
    }
}


//Anon Serviec Type Object Creation
const createModalStateObjectAnon = (modalState) => {
  console.log("creating Object modalState: " + modalState)
  return {
      modalStateAnon: modalState,
      loadingStateAnon: INIT_STATE.loadingStateAnon
  }
}

  const createLoadingStateObjectAnon= (loadingState) => {
    return {
        modalStateAnon: INIT_STATE.modalStateAnon,
        loadingStateAnon: loadingState
    }
}
