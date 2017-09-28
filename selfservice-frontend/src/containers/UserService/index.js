import React, {Component} from 'react'

import {connect} from 'react-redux';
import {doServiceStart, doFSMStart, doFSMStop,  updateFsmBasic, updateModalBasic,
        updateModalAnon, updateLoading} from '../../actions'


import Modal from 'react-modal'
import Socket  from '../../utils/socket'
import config from '../../config.json'

import "./UserService.css"


console.log(config.websocket_server);

const ip = config.websocket_server.ip;
const port = config.websocket_server.port;
const BASIC = 'basic'
const ANON = 'anon'

let url = "ws://" + ip + ":" + port + "/ws";
let ws = new WebSocket(url);
let socket = new Socket(ws);


class UserService extends Component {
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
    socket.on('fsm update', this.onFSMUpdated);
    socket.on('basic start', this.onServiceBasicStarted);
    socket.on('anon start', this.onServiceAnonStarted);
  }


//############# Socket Events ############
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

  onServiceBasicStarted = () => {

  }

  onServiceAnonStarted = () => {

  }

  onFSMStarted = (fsmData) => {
    console.log("FSM Started")
    const {updateFSMBasicAction} = this.props;
    updateFSMBasicAction(fsmData);
  }


  onFSMStopped = (fsmData) => {
    console.log("FSM Stopped")
    const {updateFSMBasicAction} = this.props;
    updateFSMBasicAction(fsmData);
  }

  onFSMUpdated =(fsmData) => {
    const {updateFSMBasicAction} = this.props;
    updateFSMBasicAction(fsmData);
  }

//############# Service Calls ############
  startService = (socket, serviceName) => {
    const {doServiceStartAction} = this.props;
    doServiceStartAction(socket, serviceName);
  }

//############# Function Calls ############
 onActionStart = (socket, fsmToStart, fsmID) => {
   const {doFSMStart} = this.props;
   doFSMStart(socket, fsmToStart, fsmID);
  }

  onActionStop =(socket, fsmToStop, fsmID) => {
    const {doFSMStop} = this.props;
    doFSMStop(socket, fsmToStop, fsmID);
  }

//############# Modal Events ############
  onUpdateLoading = (state) => {
    const {updateLoadingState} = this.props;
    updateLoadingState(state)
}

  closeModalBasic = () => {
    this.updateModalInside(false, "basic")
  }

  closeModalAnon = () => {
    this.updateModalInside(false, "anon")
  }


  updateModalInside = (state, type) => {
    console.log("Updating Modal State of Type: " + type )
    if (type === "basic"){
          const {updateModalViewBasic} = this.props;
          updateModalViewBasic(state)
    }

    if (type  === "anon") {
        const {updateModalViewAnon} = this.props;
        updateModalViewAnon(state)
    }

  }

render(){
    const customStyles = {
      content : {
        top                   : '50%',
        left                  : '50%',
        right                 : 'auto',
        bottom                : 'auto',
        marginRight           : '-50%',
        transform             : 'translate(-50%, -50%)'
      }
    };

    const {fsms} = this.props;
    const {modal} = this.props;
    console.log(fsms)
    console.log("FSMs: " + fsms.anon)
    const modalStateBasic = modal.modalStateBasic
    const loadingStateBasic = modal.loadingStateBasic

    const modalStateAnon = modal.modalStateAnon
    // const loadingStateAnon = modal.loadingStateAnon

    let theModal = null;
    if (modalStateBasic){
      theModal = <Modal
        isOpen={modalStateBasic}
        onRequestClose={this.closeModalBasic}
        style={customStyles}
        contentLabel="Basic Service Modal"
        className="modal-backdrop">
        {
          loadingStateBasic ===  false ? (
          <div className="loader"></div>
         ):(
        <div className="modal-div">
           <table className="table table-inverse">
           <thead className="thead-inverse">
            <tr>
              <th>FSM ID</th>
              <th>FSM Name</th>
              <th>State</th>
              <th>Action</th>
            </tr>
           </thead>

           <tbody>
              {fsms.basic.map((fsm) =>
                <tr key={fsm.id}  >
                  <td>{fsm.id}</td>
                  <td>{fsm.name}</td>
                  <td> {fsm.state==="started" ?
                  (<img className="img-running" src="https://d30y9cdsu7xlg0.cloudfront.net/png/22889-200.png" alt=""/>)
                : (<img className="img-running" src="https://image.flaticon.com/icons/png/512/30/30473.png" alt=""/>)
              }</td>
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
                   </td>
                </tr>
            )}
           </tbody>
           </table>

          <div className="button-modal-div">
           <button onClick={()=>this.startService(socket, BASIC)} className="btn btn-success  btn-modal d-flex justify-content-center">Start</button>
           <button onClick={()=> this.closeModalBasic()} className="btn btn-primary  btn-modal d-flex justify-content-center">close</button>
          </div>
        </div>
        )}
      </Modal>;
    }



    if (modalStateAnon){
      theModal = <Modal
        isOpen={modalStateAnon}
        onRequestClose={this.closeModalAnon}
        style={customStyles}
        contentLabel="Example Modal"
        className="modal-backdrop">
        {
          loadingStateBasic ===  false ? (
          <div className="loader"></div>
         ):(
        <div className="modal-div">
           <table className="table table-inverse">
           <thead className="thead-inverse">
            <tr>
              <th>FSM ID</th>
              <th>FSM Name</th>
              <th>State</th>
              <th>Action</th>
            </tr>
           </thead>

           <tbody>
              {fsms.anon.map((fsm) =>
                <tr key={fsm.id}  >
                  <td>{fsm.id}</td>
                  <td>{fsm.name}</td>
                  <td> {fsm.state==="started" ?
                  (<img className="img-running" src="https://d30y9cdsu7xlg0.cloudfront.net/png/22889-200.png" alt=""/>)
                : (<img className="img-running" src="https://image.flaticon.com/icons/png/512/30/30473.png" alt=""/>)
              }</td>
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
                   </td>
                </tr>
            )}
           </tbody>
           </table>
           <button onClick={()=>this.startService(socket, ANON)} className="btn btn-success  btn-modal d-flex justify-content-center">Start</button>
           <button onClick={()=>this.closeModalAnon()} className="btn btn-primary  btn-modal d-flex justify-content-center">close</button>
        </div>
        )}
      </Modal>;
    }


    return(
        <div className="container">
          {theModal}
          <div className="fullwidth">
            <div className="gallery">
              <figure className="item" onClick={() => this.updateModalInside(!modalStateBasic, "basic")}>
                <div className="img-wrap"><a><img src="https://downtownsandiego.org/wp-content/uploads/2015/01/basic-agency-logo.png" alt=""/></a></div>
                <figcaption className="caption">
                    <h3>A Basic Service That Runs Seamlinglissly</h3>

                </figcaption>
              </figure>

              <figure className="item"  onClick={() => this.updateModalInside(!modalStateAnon, "anon")}>
                <div className="img-wrap"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/a/a6/Anonymous_emblem.svg/1200px-Anonymous_emblem.svg.png" alt="" /></div>
                <figcaption className="caption">
                    <h3 className="display-3">Anonymous Routing Service</h3>
                </figcaption>
              </figure>

              <figure className="item">
                <div className="img-wrap"><img src="http://www.ptaglobal.co.uk/images/advanced_home_icon.png" alt=""/></div>
                <figcaption className="caption">
                    <h3>Advanced Services</h3>
                </figcaption>
              </figure>

              <figure className="item">
                <div className="img-wrap"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/a/a6/Anonymous_emblem.svg/1200px-Anonymous_emblem.svg.png" alt="" /></div>
                <figcaption className="caption">
                    <h3 className="display-3">Anonymous Routing Service</h3>
                </figcaption>
              </figure>
            </div>
        </div>
        </div>
      );
}
}


// const mapPropsToSto
const mapDispatchToProps = (dispatch) => ({
  doServiceStartAction: (socket, serviceName) => dispatch(doServiceStart(socket, serviceName)),
  doFSMStart:(socket, fsmToStart, fsmID)=> dispatch(doFSMStart(socket, fsmToStart, fsmID)),
  doFSMStop:(socket, fsmToStop, fsmID)=> dispatch(doFSMStop(socket, fsmToStop, fsmID)),
  updateFSMBasicAction: (fsmToUpdate) => dispatch(updateFsmBasic(fsmToUpdate)),
  updateModalViewBasic: (state, type) => dispatch(updateModalBasic(state)),
  updateModalViewAnon: (state, type) => dispatch(updateModalAnon(state)),
  updateLoadingState: (state) => dispatch(updateLoading(state))
})

const mapStateToProps = (state) => ({
  fsms: state.fsms,
  modal: state.modal
})
export default connect(mapStateToProps, mapDispatchToProps)(UserService);
