// import { browserHistory } from 'react-router';
// import {socket} from '../containers/SSMControl'
export const ITEMS_FETCH_FSMS = 'ITEMS_FETCH_FSMS'
export const UPDATE_FSMS = 'UPDATE_FSMS'
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


export const updateFsm = (fsmToUpdate) => {
  console.log("Updating:" , fsmToUpdate)
 return (dispatch) => {
   console.log(fsmToUpdate)
   dispatch({
      type: UPDATE_FSMS,
      payload: fsmToUpdate
    })
 }
}
export const doFSMStart = (socket, fsmToStart, fsmID) => {
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
