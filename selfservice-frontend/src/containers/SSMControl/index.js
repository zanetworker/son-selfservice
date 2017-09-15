import React, {Component} from 'react'
import {connect} from 'react-redux';
import "./SSMControl.css"
import {doFSMStart, doFSMStop,  updateFsm} from '../../actions'
import Socket  from '../../utils/socket'
import config from '../../config.json'

console.log(config.websocket_server);

const ip = config.websocket_server.ip;
const port = config.websocket_server.port;

let url = "ws://" + ip + ":" + port + "/ws";
let ws = new WebSocket(url);
let socket = new Socket(ws);


class SSMControl extends Component {
  constructor(props){
    super(props);
    this.state = {
      connected: false,
      nameToStart: "",
    }
  }

  handleClick = (name) => {
    console.log(name)
  }

  componentDidMount(){
    socket.on('connect', this.onConnect);
    socket.on('disconnect', this.onDisconnect);
    socket.on('fsm start', this.onFSMStarted);
    socket.on('fsm stop', this.onFSMStopped);
    socket.on('fsm update', this.onFSMUpdated)
  }

  onConnect = () => {
    this.setState({
      connected: true
    });
  }

  onDisconnect = () => {
    this.setState({
      connected: false
    });
  }

  onFSMStarted = (fsmData) => {
    console.log("FSM Started")
    const {updateFSMAction} = this.props;
    updateFSMAction(fsmData);
  }


  onFSMStopped = (fsmData) => {
    console.log("FSM Stopped")
    const {updateFSMAction} = this.props;
    updateFSMAction(fsmData);
  }

  onFSMUpdated =(fsmData) => {
    const {updateFSMAction} = this.props;
    updateFSMAction(fsmData);
  }

 onActionStart = (socket, fsmToStart, fsmID) => {
   const {doFSMStart} = this.props;
   doFSMStart(socket, fsmToStart, fsmID);
  }

  onActionStop =(socket, fsmToStop, fsmID) => {
    const {doFSMStop} = this.props;
    doFSMStop(socket, fsmToStop, fsmID);
  }

  onActionResume =(id) => {
    //TODO
    console.log(id)
  }

render(){
  const {fsms} = this.props;
  return(
    <div>
      <legend> SSM Information & Control </legend>

    <table className="table">
   <thead className="thead-inverse">
     <tr>
       <th>FSM ID</th>
       <th>FSM Name</th>
       <th>State</th>
       <th>Action</th>
     </tr>
   </thead>

   <tbody>
       {fsms.map((fsm) =>
         <tr key={fsm.id}  >
           <td>{fsm.id}</td>
           <td>{fsm.name}</td>
           <td>{fsm.state}</td>
           <td className="buttons-sep insert-margin">
           {
             fsm.state ===  "started" ? (
            <button type="button" className="btn btn-success disabled" onClick={() => this.onActionStart(socket, fsm.name, fsm.id)}>Start</button>
            ):(
            <button  type="button" className="btn btn-success" onClick={() => this.onActionStart(socket, fsm.name, fsm.id)}>Start</button>
           )}

            <span></span><span></span>
            {fsm.state ===  "stopped" ? (
              <button type="button" className="btn btn-danger disabled" onClick={() => this.onActionStop(socket, fsm.name, fsm.id)}>Stop</button>
             ):(
             <button  type="button" className="btn btn-danger" onClick={() => this.onActionStop(socket, fsm.name, fsm.id)}>Stop</button>
            )}
            <span></span><span></span>
            <button  type="button"className="btn btn-primary" onClick={() => this.onActionResume(fsm.id)}>Resume</button>
            </td>
         </tr>
     )}
   </tbody>
 </table>
  </div>
  );
}
}


// const mapPropsToSto
const mapDispatchToProps = (dispatch) => ({
  doFSMStart:(socket, fsmToStart, fsmID)=> dispatch(doFSMStart(socket, fsmToStart, fsmID)),
  doFSMStop:(socket, fsmToStop, fsmID)=> dispatch(doFSMStop(socket, fsmToStop, fsmID)),
  updateFSMAction: (fsmToUpdate) => dispatch(updateFsm(fsmToUpdate)),
})

const mapStateToProps = (state) => ({
  fsms: state.fsms
})
export default connect(mapStateToProps, mapDispatchToProps)(SSMControl);
