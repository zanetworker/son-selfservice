import {UPDATE_FSMS} from '../actions'

const INIT_STATE = [
  // {
  //     id: "12345",
  //     name: "Firewall",
  //     state: "stopped"
  // },
  // {
  //     id: "2",
  //     name: "Loadbalancer",
  //     state: "started"
  // }
];



export default (state = INIT_STATE, action) => {
  switch (action.type){
    case UPDATE_FSMS:
      const stateToReturn = state.filter((fsm)=>{
        return fsm.id !== action.payload.id
      })
      stateToReturn.push(createStateObject(action.payload))
      return stateToReturn
    default:
      return state;
  }
}


const createStateObject= (dataForObject) => {
  return {
      id: dataForObject.id,
      name: "Firewall",
      state: dataForObject.state
  }
}
