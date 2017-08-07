const INIT_STATE = [
  {
      id: 1,
      name: "Firewall",
      state: "stopped"
  },
  {
      id: 2,
      name: "Loadbalancer",
      state: "started"
  }];



export default (state = INIT_STATE, action) => {
  switch (action.type){
    default:
      return state;
  }
}
