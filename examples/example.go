package main
import "github.com/bigstepinc/metal-cloud-sdk-go"
import "os"
import "log"

func main(){
  user := os.Getenv("METALCLOUD_USER")
  apiKey := os.Getenv("METALCLOUD_API_KEY")
  endpoint := os.Getenv("METALCLOUD_ENDPOINT")

  if(user=="" || apiKey=="" || endpoint==""){
    log.Fatal("METALCLOUD_USER, METALCLOUD_API_KEY, METALCLOUD_ENDPOINT environment variables must be set")
  }

  client, err := metalcloud.GetMetalcloudClient(user, apiKey, endpoint)
  if err != nil {
    log.Fatal("Error initiating client: %s", err)
  }

  infras,err :=client.Infrastructures()
  if err != nil {
    log.Fatal("Error retrieving a list of infrastructures: %s", err)
  }

  for _,infra := range *infras{
    log.Printf("%s(%d)",infra.InfrastructureLabel, infra.InfrastructureID)
  }


}