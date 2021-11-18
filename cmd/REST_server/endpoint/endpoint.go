package main

import "github.com/neidersalgado/go-camp-grpc/cmd/REST_server/bussiness/usecase"


func makeDeleteUserEndpoint(s usecase.UserService) endpoint.Endpoint { 
	return func(ctx context.Context, request interface{})  
  		(interface{}, error) {  
      req := request.(DeleteUserRequest)
      msg, err := s.Delete(req.UserID)
    return DeleteUserResponse{Msg: msg, Err: err}, nil }
}

func makeCreateUserEndpoint(s usecase.UserService) endpoint.Endpoint { 
	return func(ctx context.Context, request interface{})  
  		(interface{}, error) {  
      req := request.(CreateUserRequest)
	    model := EntityToModel(req.User)
      msg, err := s.Create(model)
    return CreateUserResponse{Msg: msg, Err: err}, nil }
}



func makeGetUserEndpoint(s usecase.UserService) endpoint.Endpoint { 
	return func(ctx context.Context, request interface{})  
  		(interface{}, error) {  
      req := request.(GetUserRequest)
      entity, err := s.GetByID(red.UserID)
	  model := ModelToEntity(entity)
    return GetUserResponse{User: model, Err: err}, nil }
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) { 
  var req CreateUserRequest 
  if err := json.NewDecoder(r.Body).Decode(&req.user); err !=        nil 
   {  
      return nil, err 
   }
   return req, nil
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) { 
  var req GetUserRequest 
  vars := mux.Vars(r) 
  req = GetUserRequest{ UserID: vars["user_id"], } 
  return req, nil
}

func decodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) { 
  var req DeleteUserRequest 
  vars := mux.Vars(r) 
  req = DeleteUserRequest{ UserID : vars["user_id"], } 
  return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error 
{ 
  w.Header().Set("Content-Type", "application/json; charset=utf-8")     
  return json.NewEncoder(w).Encode(response)
}