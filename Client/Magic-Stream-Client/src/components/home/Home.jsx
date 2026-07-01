import {useState, useEffect} from 'react';
import axiosClient from '../../api/axiosConfig'
import Movies from '../movies/Movies';
import Spinner from '../spinner/Spinner';
import Hero from '../hero/Hero';

const Home =() => {

    const [movies, setMovies] = useState([]);
    const [loading, setLoading] = useState(false)
    const [message, setMessage] = useState();
    const [search,setSearch] = useState("");
    const [searchTerm, setSearchTerm] = useState("");

    useEffect(() => {
        
        const fetchMovies = async () => {
            setLoading(true);
            setMessage("");
            try{
                const response = await axiosClient.get('/movies');
                
                console.log("Full Response:", response);
                console.log("Movies Data:", response.data);
                
                setMovies(response.data);
                if (response.data.length === 0){
                    setMessage('There are currently no movies available')
                }

            }catch(error){
                console.error('Error fetching movies:', error)
                setMessage("Error fetching movies")
            }finally{
                setLoading(false)
            }
        }
        fetchMovies();
    }, []);

    const filteredMovies =
            movies.filter((movie)=>
               movie.title.toLowerCase().includes(search.toLowerCase())
            );

    return (
        <>
            <Hero />

            <div className="container mt-4">
                <input
                    type="text"
                    className="form-control form-control-lg"
                    placeholder="🔍 Search movies..."
                    value={search}
                    onChange={(e) => setSearch(e.target.value)}
                />
            </div>

            <div className="container mt-5">
    <h2 className="text-light fw-bold">
        Trending Movies
    </h2>
</div>

            {loading ? (
                <Spinner />
            ) : (
                <Movies
                    movies={filteredMovies}
                    message={message}
                />
            )}
        </>
    );

}

export default Home;