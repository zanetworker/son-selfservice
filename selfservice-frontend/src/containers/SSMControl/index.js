import React, {Component} from 'react'
import {connect} from 'react-redux';
import "./SSMControl.css"
import {fetchFsms} from '../../actions'
import Socket  from '../../utils/socket'


let url = "ws://127.0.0.1:4000/wsecho";
let ws = new WebSocket(url);
let socket = new Socket(ws);

export {socket};

class SSMControl extends Component {

  constructor(props){
    super(props);
    this.state = {
      connected: false,
    }
  }

  componentDidMount(){
    socket.on('connect', this.onConnect);
    socket.on('disconnect', this.onDisconnect);
    socket.on('fsm add', this.onAddFSM);
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

  onAddFSM = (fsm) => {
    console.log("Adding FSM");
    console.log(fsm);
  }

 onActionStart =(socket) => {
   const {itemFetchFsmsAction} = this.props;
  itemFetchFsmsAction(socket)
  }

  onActionStop =(id) => {
    console.log(id);
  }

  onActionResume =(id) => {
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
           {fsm.state ===  "started" ? (
            <button type="button" className="btn btn-success disabled" onClick={() => this.onActionStart(this.state.socket)}>Start</button>
            ):(
            <button  type="button" className="btn btn-success" onClick={() => this.onActionStart(socket)}>Start</button>
           )}

            <span></span><span></span>
            {fsm.state ===  "stopped" ? (
              <button type="button" className="btn btn-danger disabled" onClick={() => this.onActionStop(fsm.id)}>Stop</button>
             ):(
             <button  type="button" className="btn btn-danger" onClick={() => this.onActionStop(fsm.id)}>Stop</button>
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
  itemFetchFsmsAction:(socket)=> dispatch(fetchFsms(socket)),
})

const mapStateToProps = (state) => ({
  fsms: state.fsms
})
export default connect(mapStateToProps, mapDispatchToProps)(SSMControl);
