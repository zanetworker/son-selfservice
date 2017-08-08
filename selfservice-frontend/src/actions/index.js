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
  console.log("Updateing:" , fsmToUpdate)
 return (dispatch) => {
   dispatch({
      type: UPDATE_FSMS,
      payload: fsmToUpdate
    })
 }

}
export const fetchFsms = (socket) => {
  return (dispatch) => {
    let msg = {
      name: 'fsm start',
      data: {
        id: '12345',
        action: 'fsm_name'
      }
    }

    console.log(msg);
    socket.emit(msg)
    console.log("message sent")
    // fetch(url)
    //   .then((response) => response.json())
    //   .then((fsms) => dispatch(fsms))
    //   .catch(()=> console.log("failed"))
  };
}
