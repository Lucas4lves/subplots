export type Plot = {
    id : number
    title : string 
    story : string 
    status : boolean
    created_at : string
    updated_at : string 
}

export const Plot = ({id, title, story, status, created_at, updated_at } : Plot) => {

    let idStr = String(id)

    return (
    <>
        <div id={idStr} className="plot-card-body">
            <h2 className="plot-card-title">
                {title}
            </h2>

            <p className="plot-card-story">
                {story}
            </p>

            <p className="plot-card-status">{status ? "Done" : "On going"}</p>
            <p className="plot-card-created-at">{created_at}</p>
            <p className="plot-card-updated-at">{updated_at}</p>
        </div>
    </>)
}