// import { browserHistory } from 'react-router';
// import {socket} from '../containers/SSMControl'
export const ITEMS_FETCH_FSMS = 'ITEMS_FETCH_FSMS'
export const UPDATE_FSMS_BASIC= 'UPDATE_FSMS_BASIC'
export const UPDATE_FSMS_ANON= 'UPDATE_FSMS_ANON'
export const UPDATE_MODAL = 'UPDATE_MODAL'
export const UPDATE_LOADING = 'UPDATE_LOADING'
// export const itemsFetchData = (url) => {
//
//   return (dispatch) => {
//     fetch(url)
//         .then((response) => {
//             if (!response.ok) {
//                 throw Error(response.statusText);
//             }
//
//             dispatch(itemsIsLoading(false));
//
//             return response;
//         })
//         .then((response) => response.json())
//         .then((items) => dispatch(itemsFetchDataSuccess(items)))
//         .catch(() => dispatch(itemsHasErrored(true)));
//     browserHistory.push(`/login`)
//   };
// }

export const updateLoading= (state) => {
  console.log("Loading State: " + state)
  return (dispatch) =>{
  let msg = {
       type: UPDATE_LOADING,
       payload: state
  }
  dispatch(msg)
  };
}


export const updateModalBasic= (state) => {
  console.log("[actions] updateModalBasic")
  return (dispatch) =>{
        let msgBasic = {
           type: UPDATE_MODAL,
           payload: {
             state: state,
             type: "basic"
           }
         }
        dispatch(msgBasic)
}
}

export const updateModalAnon = (state) => {
  console.log("[actions] updateModalAnon")
  return (dispatch) => {
        let msgAnon = {
           type: UPDATE_MODAL,
           payload: {
             state: state,
             type: "anon"
           }
          }
       dispatch(msgAnon)
    }
  }


export const updateFsmBasic = (fsmToUpdate) => {
  console.log("Updating Basic:" , fsmToUpdate)
 return (dispatch) => {
   console.log(fsmToUpdate)
   dispatch({
      type: UPDATE_FSMS_BASIC,
      payload: fsmToUpdate
    })
 };
}

export const updateFsmAnon = (fsmToUpdate) => {
console.log("Updating Anon:" , fsmToUpdate)
 return (dispatch) => {
   console.log(fsmToUpdate)
   dispatch({
      type: UPDATE_FSMS_ANON,
      payload: fsmToUpdate
    })
 };
}


export const doServiceStart = (socket, serviceType) => {
  console.log("hello: " + serviceType)
  console.log("Starting service: " + serviceType + "...!")
  return (dispatch) => {
    let msg = null;
    console.log("stuff")
    if (serviceType === "basic"){
      msg = {
        name: 'basic start',
        data: ""
      }
    }

    if (serviceType === "anon"){
        msg = {
        name: 'anon start',
        data: ""
      }
    }
    socket.emit(msg)

  };
}


export const doFSMStart = (socket, fsmToStart, fsmID) => {
  console.log("Starting FSM...!")
  return (dispatch) => {
    let msg = {
      name: 'fsm start',
      data: {
        id: fsmID,
        name: fsmToStart
      }
    }

    socket.emit(msg)
    console.log("Sent Start FSM Command!")
     // fetch(url)
    //   .then((response) => response.json())
    //   .then((fsms) => dispatch(fsms))
    //   .catch(()=> console.log("failed"))
  };
}


export const doFSMStop = (socket, fsmToStop, fsmID) => {
  return (dispatch) => {
    let msg = {
      name: 'fsm stop',
      data: {
        id: fsmID,
        name: fsmToStop
      }
    }

    socket.emit(msg)
    console.log("Sent Stop FSM Command!")
    // fetch(url)
    //   .then((response) => response.json())
    //   .then((fsms) => dispatch(fsms))
    //   .catch(()=> console.log("failed"))
  };
}
