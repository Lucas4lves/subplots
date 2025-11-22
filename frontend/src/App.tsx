import {Plot} from "../components/Plot.tsx"

const samplePlot : Plot = {
  id: 1,
  title: "Learn React",
  story: "Getting up to speed with basics",
  status: true,
  created_at: new Date().toDateString(),
  updated_at: new Date().toDateString()
}
function App() {
  
  return (
    <>
      <Plot {...samplePlot}/>
    </>
  )
}

export default App
