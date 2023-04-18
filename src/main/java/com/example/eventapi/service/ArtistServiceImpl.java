package com.example.eventapi.service;

import com.example.eventapi.exception.CustomWebClientException;
import com.example.eventapi.models.Artist;
import com.example.eventapi.models.ArtistInfo;
import com.example.eventapi.models.Event;
import com.example.eventapi.models.Venue;
import com.example.eventapi.repository.ArtistRepository;
import com.example.eventapi.repository.EventRepository;
import com.example.eventapi.repository.VenueRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;
import reactor.core.publisher.Mono;

import java.util.Collections;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import java.util.stream.Collectors;

@Service
public class ArtistServiceImpl implements  ArtistService{

    private final ArtistRepository artistRepository;
    private final EventRepository eventRepository;
    private final VenueRepository venueRepository;

    @Autowired
    public ArtistServiceImpl(ArtistRepository artistRepository, EventRepository eventRepository, VenueRepository venueRepository) {
        this.artistRepository = artistRepository;
        this.eventRepository = eventRepository;
        this.venueRepository = venueRepository;
    }

    @Override
    public Mono<ArtistInfo> getArtistInfo(String artistId) {
        Mono<List<Artist>> artistsMono = artistRepository.fetchArtists();

        Mono<List<Event>> eventsMono = eventRepository.fetchEvents();

        Mono<List<Venue>> venuesMono = venueRepository.fetchVenues();

        return artistsMono
                .flatMap(artists -> {
                    Optional<Artist> artistOpt = artists.stream().filter(a -> a.getId().equals(artistId)).findFirst();
                    if (!artistOpt.isPresent()) {
                        String errorMessage = "Artist not found for artistId " + artistId;
                        return Mono.error(new CustomWebClientException(errorMessage, HttpStatus.NOT_FOUND));
                    }
                    Artist artist = artistOpt.get();
                    return Mono.zip(eventsMono, venuesMono, (events, venues) -> {
                        List<Event> artistEvents = events.stream()
                                .filter(event -> event.getArtists().stream().anyMatch(a -> a.getId().equals(artistId)))
                                .peek(event -> {
                                    Optional<Venue> venue = venues.stream()
                                            .filter(v -> v.getId().equals(event.getVenue().getId()))
                                            .findFirst();
                                    venue.ifPresent(event::setVenue);
                                })
                                .collect(Collectors.toList());
                        return ArtistInfo.builder().artist(artist).events(artistEvents).build();
                    });
                });
    }
}