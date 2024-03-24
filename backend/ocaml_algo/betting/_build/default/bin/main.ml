let concate (stringList: string list) (connector:string) :string = 
  
  let rec concatHelper (stringList: string list) (connector:string) (acc:string) = 
    match stringList with
    |[]-> print_string acc
    |[tl] -> acc^tl
    |hd::tl-> concatHelper tl connector (acc^hd^connector)
                
  in
  concatHelper stringList connector ""
;;

concate ["aa"; "bb"; "cc"] "->";;
concate ["aa"] "->";;
concate [] "->";;
