package controller

import (
	"GoProject/token"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"GoProject/db"	
)

func LoginTeacher(w http.ResponseWriter, r *http.Request){
	reqBody, _ := ioutil.ReadAll(r.Body)
	 var teacher db.Teacher
	 json.Unmarshal(reqBody,&teacher)
	 findTeacher:=db.LoginTeacher(teacher.Email)
	 if findTeacher.Id==0{
		json.NewEncoder(w).Encode(Response{Status:404,Msg:"User Not Exists"})
	 }else if findTeacher.Password!=teacher.Password{
		json.NewEncoder(w).Encode(Response{Status:401,Msg:"Unauthroized Worng Password"})

	 }else{
		 token,err:=token.GenerateJWT(findTeacher.Email)
		 if err!=nil{
			 json.NewEncoder(w).Encode(Response{Status:404,Msg:"Some Error Occured"})
		return
		 }
		json.NewEncoder(w).Encode(User{Token:token,User:findTeacher})
	 }
	 
}
