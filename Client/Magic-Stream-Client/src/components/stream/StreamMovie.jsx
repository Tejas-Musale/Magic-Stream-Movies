import {useParams} from 'react-router-dom';
import ReactPlayer from 'react-player';
import './StreamMovie.css';

// const StreamMovie = () => {
//     const { yt_id } = useParams();

//     console.log("YT ID:", yt_id);

//     return (
//         <div className="container mt-4">
//             <ReactPlayer
//                 url={`https://www.youtube.com/watch?v=${yt_id}`}
//                 controls
//                 width="100%"
//                 height="100%"
//             />
//         </div>
//     );
// };

// export default StreamMovie;

const StreamMovie = () => {
    const { yt_id } = useParams();

    return (
        <div className="container mt-4">
            <ReactPlayer
                src={`https://www.youtube.com/watch?v=${yt_id}`}
                controls
                width="100%"
                height="600px"
            />
        </div>
    );
};


export default StreamMovie;