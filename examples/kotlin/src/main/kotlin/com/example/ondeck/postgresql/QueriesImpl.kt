// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package com.example.ondeck.postgresql

import java.sql.Connection
import java.sql.SQLException
import java.sql.Statement
import java.sql.Types
import java.time.LocalDateTime

const val createCity = """-- name: createCity :one
INSERT INTO city (
    name,
    slug
) VALUES (
    ?,
    ?
) RETURNING slug, name
"""

const val createVenue = """-- name: createVenue :one
INSERT INTO venue (
    slug,
    name,
    city,
    created_at,
    spotify_playlist,
    status,
    statuses,
    tags
) VALUES (
    ?,
    ?,
    ?,
    NOW(),
    ?,
    ?,
    ?,
    ?
) RETURNING id
"""

const val deleteVenue = """-- name: deleteVenue :exec
DELETE FROM venue
WHERE slug = ? AND slug = ?
"""

const val getCity = """-- name: getCity :one
SELECT slug, name
FROM city
WHERE slug = ?
"""

const val getVenue = """-- name: getVenue :one
SELECT id, status, statuses, slug, name, city, spotify_playlist, songkick_id, tags, created_at
FROM venue
WHERE slug = ? AND city = ?
"""

const val listCities = """-- name: listCities :many
SELECT slug, name
FROM city
ORDER BY name
"""

const val listVenues = """-- name: listVenues :many
SELECT id, status, statuses, slug, name, city, spotify_playlist, songkick_id, tags, created_at
FROM venue
WHERE city = ?
ORDER BY name
"""

const val updateCityName = """-- name: updateCityName :exec
UPDATE city
SET name = ?
WHERE slug = ?
"""

const val updateVenueName = """-- name: updateVenueName :one
UPDATE venue
SET name = ?
WHERE slug = ?
RETURNING id
"""

const val venueCountByCity = """-- name: venueCountByCity :many
SELECT
    city,
    count(*)
FROM venue
GROUP BY 1
ORDER BY 1
"""

data class VenueCountByCityRow (
  val city: String,
  val count: Long
)

class QueriesImpl(private val conn: Connection) : Queries {

// Create a new city. The slug must be unique.
// This is the second line of the comment
// This is the third line

  @Throws(SQLException::class)
  override fun createCity(name: String, slug: String): City? {
    return conn.prepareStatement(createCity).use { stmt ->
      stmt.setString(1, name)
          stmt.setString(2, slug)

      val results = stmt.executeQuery()
      if (!results.next()) {
        return null
      }
      val ret = City(
                results.getString(1),
                results.getString(2)
            )
      if (results.next()) {
          throw SQLException("expected one row in result set, but got many")
      }
      ret
    }
  }

  @Throws(SQLException::class)
  override fun createVenue(
      slug: String,
      name: String,
      city: String,
      spotifyPlaylist: String,
      status: Status,
      statuses: List<Status>,
      tags: List<String>): Int? {
    return conn.prepareStatement(createVenue).use { stmt ->
      stmt.setString(1, slug)
          stmt.setString(2, name)
          stmt.setString(3, city)
          stmt.setString(4, spotifyPlaylist)
          stmt.setObject(5, status.value, Types.OTHER)
          stmt.setArray(6, conn.createArrayOf("status", statuses.map { v -> v.value }.toTypedArray()))
          stmt.setArray(7, conn.createArrayOf("text", tags.toTypedArray()))

      val results = stmt.executeQuery()
      if (!results.next()) {
        return null
      }
      val ret = results.getInt(1)
      if (results.next()) {
          throw SQLException("expected one row in result set, but got many")
      }
      ret
    }
  }

  @Throws(SQLException::class)
  override fun deleteVenue(slug: String) {
    conn.prepareStatement(deleteVenue).use { stmt ->
      stmt.setString(1, slug)
          stmt.setString(2, slug)

      stmt.execute()
    }
  }

  @Throws(SQLException::class)
  override fun getCity(slug: String): City? {
    return conn.prepareStatement(getCity).use { stmt ->
      stmt.setString(1, slug)

      val results = stmt.executeQuery()
      if (!results.next()) {
        return null
      }
      val ret = City(
                results.getString(1),
                results.getString(2)
            )
      if (results.next()) {
          throw SQLException("expected one row in result set, but got many")
      }
      ret
    }
  }

  @Throws(SQLException::class)
  override fun getVenue(slug: String, city: String): Venue? {
    return conn.prepareStatement(getVenue).use { stmt ->
      stmt.setString(1, slug)
          stmt.setString(2, city)

      val results = stmt.executeQuery()
      if (!results.next()) {
        return null
      }
      val ret = Venue(
                results.getInt(1),
                Status.lookup(results.getString(2))!!,
                (results.getArray(3).array as Array<String>).map { v -> Status.lookup(v)!! }.toList(),
                results.getString(4),
                results.getString(5),
                results.getString(6),
                results.getString(7),
                results.getString(8),
                (results.getArray(9).array as Array<String>).toList(),
                results.getObject(10, LocalDateTime::class.java)
            )
      if (results.next()) {
          throw SQLException("expected one row in result set, but got many")
      }
      ret
    }
  }

  @Throws(SQLException::class)
  override fun listCities(): List<City> {
    return conn.prepareStatement(listCities).use { stmt ->
      
      val results = stmt.executeQuery()
      val ret = mutableListOf<City>()
      while (results.next()) {
          ret.add(City(
                results.getString(1),
                results.getString(2)
            ))
      }
      ret
    }
  }

  @Throws(SQLException::class)
  override fun listVenues(city: String): List<Venue> {
    return conn.prepareStatement(listVenues).use { stmt ->
      stmt.setString(1, city)

      val results = stmt.executeQuery()
      val ret = mutableListOf<Venue>()
      while (results.next()) {
          ret.add(Venue(
                results.getInt(1),
                Status.lookup(results.getString(2))!!,
                (results.getArray(3).array as Array<String>).map { v -> Status.lookup(v)!! }.toList(),
                results.getString(4),
                results.getString(5),
                results.getString(6),
                results.getString(7),
                results.getString(8),
                (results.getArray(9).array as Array<String>).toList(),
                results.getObject(10, LocalDateTime::class.java)
            ))
      }
      ret
    }
  }

  @Throws(SQLException::class)
  override fun updateCityName(name: String, slug: String) {
    conn.prepareStatement(updateCityName).use { stmt ->
      stmt.setString(1, name)
          stmt.setString(2, slug)

      stmt.execute()
    }
  }

  @Throws(SQLException::class)
  override fun updateVenueName(name: String, slug: String): Int? {
    return conn.prepareStatement(updateVenueName).use { stmt ->
      stmt.setString(1, name)
          stmt.setString(2, slug)

      val results = stmt.executeQuery()
      if (!results.next()) {
        return null
      }
      val ret = results.getInt(1)
      if (results.next()) {
          throw SQLException("expected one row in result set, but got many")
      }
      ret
    }
  }

  @Throws(SQLException::class)
  override fun venueCountByCity(): List<VenueCountByCityRow> {
    return conn.prepareStatement(venueCountByCity).use { stmt ->
      
      val results = stmt.executeQuery()
      val ret = mutableListOf<VenueCountByCityRow>()
      while (results.next()) {
          ret.add(VenueCountByCityRow(
                results.getString(1),
                results.getLong(2)
            ))
      }
      ret
    }
  }

}

