import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import Select from '@material-ui/core/Select';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import { EntUser } from '../../api/models/EntUser';
import { EntPatient } from '../../api/models/EntPatient';
import { EntStatusTracking } from '../../api/models/EntStatusTracking';

const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
     margin: theme.spacing(3),
   },
   withoutLabel: {
     marginTop: theme.spacing(3),
   },
   textField: {
     width: '25ch',
   },
 }),
);

//const username={ givenuser: 'นพ.เอื้อสีวงศ์ มะม่วงร่วง'}
export default function Create() {
 const classes = useStyles();
 const profile = { givenName: 'ระบบประวัติการรักษา' };
 const api = new DefaultApi();
 
 //const [user, setUser] = useState(initialUserState);
 const [users, setUsers] = useState<EntUser[]>([]);
 const [patients, setPatients] = useState<EntPatient[]>([]);
 const [statuss, setStatuss] = useState<EntStatusTracking[]>([]);
 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);
 const [loading, setLoading] = useState(true);

 //const [treatment, setTreatment] = useState(String);
 const [statustrack, setStatustrack] = useState(Number);
 const [symptom, setSymptom] = useState(String);
 const [userid, setUserid] = useState(Number);
 const [patientid, setPatientid] = useState(Number);
 const [addtime, setAddtime] = useState(String);
 
 useEffect(() => {
   const getUsers = async () => {
     const u = await api.listUser({ limit: 10, offset: 0 });
     setLoading(false);
     setUsers(u);
   };
   getUsers();

   const getPatients = async () => {
    const p = await api.listPatient({ limit: 10, offset: 0 });
    setLoading(false);
    setPatients(p);
  };
  getPatients();

  const getStatuss = async () => {
    const st = await api.listStatustracking({ limit: 10, offset: 0 });
    setLoading(false);
    setStatuss(st);
  };
  getStatuss();

 }, [loading]);
 
 const handlesymptomChange = (event: any) => {
  setSymptom(event.target.value as string);
 };

 const handletimeChange = (event: any) => {
  setAddtime(event.target.value as string);
 };

 React.useEffect(() => {
   const data = localStorage.getItem("doctordata");
   if (data) {
     setUserid(JSON.parse(data));
   }
 })

 const CreateTreatment = async () =>{
  const treatment = {
    userid    : userid,
	  patientid : patientid,
	  status    : statustrack,
	  symptom   : symptom,
    added     : addtime + "T00:00:00Z"
  }

   const res:any = await api.createTreatment({ treatment : treatment });
   setStatus(true);
   if (res.id != ''){
     setAlert(true);
   } else {
     setAlert(false);
   }
   const timer = setTimeout(() => {
     setStatus(false);
   }, 1000);
 };


//  const doctor_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
//   setUserid(event.target.value as number);
//  };

const patient_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setPatientid(event.target.value as number);
 };

 const status_patienthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setStatustrack(event.target.value as number);
 };
 
 return (
   <Page theme={pageTheme.service}>
     <Header
       title={`Welcome ${profile.givenName || 'ระบบประว้ติ'}`}
     //  subtitle = {`${username.givenuser}`}
     ></Header>
     <Content>
       <ContentHeader title="ระบบลงทะเบียนประวัติการรักษา">
{/* 
       <Button variant="outlined" color="secondary" href="#outlined-buttons">
          LogOut
        </Button>
*/}
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                ✔️ ADD DATA COMPLETE : 😁 😆 🤩!
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 ❌ CAN'T ADD DATA : 🥺 😵 😱!
               </Alert>
             )}
           </div>
         ) : null}
       </ContentHeader>

       <div className={classes.root}>
         <form noValidate autoComplete="off">
         <div>
<FormControl 
             className={classes.margin}
             variant="outlined"
           >
            <InputLabel id="doctor_id-label">DoctorName</InputLabel>
            <Select
              labelId="doctor_id-label"
              label="Doctor"
              id="doctor_id"
              value={userid}
            //  onChange={doctor_id_handleChange}
              style = {{width: 600}}
            >
              {users.map((item:EntUser)=>
              <MenuItem value={item.id}>{item.doctorName}</MenuItem>)}
            </Select>
 </FormControl>
 </div>
<div>
<FormControl 
             className={classes.margin}
             variant="outlined"
           >
            <InputLabel id="patient_id-label">PatientName</InputLabel>
            <Select
              labelId="patient_id-label"
              label="Patient"
              id="patient_id"
              value={patientid}
              onChange={patient_id_handleChange}
              style = {{width: 600}}
            >
              {patients.map((item:EntPatient)=>
              <MenuItem value={item.id}>{item.patientName}</MenuItem>)}
            </Select>
 </FormControl>
 </div>

           <FormControl
            // fullWidth
             className={classes.margin}
             variant="outlined"
           >
             <TextField
               id="symptom"
               label="Symptom"
               variant="outlined"
               type="string"
               size="medium"
               value={symptom}
               onChange={handlesymptomChange}
               style = {{width: 600}}
             />
           </FormControl>

<div>
<FormControl 
             className={classes.margin}
             variant="outlined"
           >
            <InputLabel id="status-label">Status</InputLabel>
            <Select
              labelId="status-label"
              label="Stauts"
              id="status_patient"
              value={statustrack}
              onChange={status_patienthandleChange}
              style = {{width: 600}}
            >
              {statuss.map((item:EntStatusTracking)=>
              <MenuItem value={item.id}>{item.status}</MenuItem>)}
            </Select>
          </FormControl>
 </div>
<table>
  <tr><td>
  📅 หากรักษาเสร็จสิ้นให้เลือก " วันนี้ "
    <p>
  📅 หากยังติดตามอาการให้เลือกวันที่นัดหมาย
    </p>
    </td><td>
 <FormControl className={classes.margin} >
      <TextField
        id="Appoint_date"
        label="Appoint_date"
        type="date"
        value={addtime}
        onChange={handletimeChange}
       // defaultValue={addtime.now}
        className={classes.textField}
        InputLabelProps={{
          shrink: true,
        }}
        
      />
    </FormControl>
    </td></tr></table>

           <div className={classes.margin}>
             <Button
               onClick={() => {
                 CreateTreatment();
               }}
               variant="contained"
               color="primary"
             >
               Confirm
             </Button>
             <Button
               style={{ marginLeft: 20 }}
               component={RouterLink}
               to="/welcome"
               variant="contained"
             >
               Back
             </Button>
           </div>
           
         </form>
       </div>
     </Content>
   </Page>
 );
}
